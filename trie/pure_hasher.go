// Copyright 2019 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package trie

import (
	"sync"

	"github.com/ethereum/go-ethereum/rlp"
	"golang.org/x/crypto/sha3"
)

// pureHasher is a type used for the trie Hash operation. A pureHasher has some
// internal preallocated temp space
type pureHasher struct {
	sha keccakState

	tmp    sliceBuffer
	tmpKey []byte
}

// pureHasherPool holds pureHashers
var pureHasherPool = sync.Pool{
	New: func() interface{} {
		return &pureHasher{
			tmp:    make(sliceBuffer, 0, 550), // cap is as large as a full fullNode.
			tmpKey: make([]byte, 64),          // space for an packed key
			sha:    sha3.NewLegacyKeccak256().(keccakState),
		}
	},
}

func newPureHasher() *pureHasher {
	h := pureHasherPool.Get().(*pureHasher)
	return h
}

func returnPureHasherToPool(h *pureHasher) {
	pureHasherPool.Put(h)
}

// hash collapses a node down into a hash node, also returning a copy of the
// original node initialized with the computed hash to replace the original one.
func (h *pureHasher) hash(n node, force bool) (hashed node, cached node) {
	// We're not storing the node, just hashing, use available cached data
	if hash, _ := n.cache(); hash != nil {
		return hash, n
	}
	// Trie not processed yet or needs storage, walk the children
	switch n := n.(type) {
	case *shortNode:
		collapsed, cached := h.hashShortNodeChildren(n)
		hashed := h.shortnodeToHash(collapsed, force)
		// We need to retain the possibly _not_ hashed node, in case it was too
		// small to be hashed
		if hn, ok := hashed.(hashNode); ok {
			cached.flags.hash = hn
		} else {
			cached.flags.hash = nil
		}
		return hashed, cached
	case *fullNode:
		collapsed, cached := h.hashFullNodeChildren(n)
		hashed = h.fullnodeToHash(collapsed, force)
		if hn, ok := hashed.(hashNode); ok {
			cached.flags.hash = hn
		} else {
			cached.flags.hash = nil
		}
		return hashed, cached
	default:
		// Value and hash nodes don't have children so they're left as were
		return n, n
	}
}

// hashShortNodeChildren collapses the short node. The returned collapsed node
// holds a live reference to the Key, and must not be modified.
// The cached
func (h *pureHasher) hashShortNodeChildren(n *shortNode) (collapsed, cached *shortNode) {
	// Hash the short node's child, caching the newly hashed subtree
	collapsed, cached = n.copy(), n.copy()
	// Previously, we did copy this one. We don't seem to need to actually
	// do that, since we don't overwrite/reuse keys
	//cached.Key = common.CopyBytes(n.Key)
	collapsed.Key = hexToCompact(n.Key)
	// Unless the child is a valuenode or hashnode, hash it
	switch n.Val.(type) {
	case *fullNode, *shortNode:
		collapsed.Val, cached.Val = h.hash(n.Val, false)
	}
	return collapsed, cached
}

func (h *pureHasher) hashFullNodeChildren(n *fullNode) (collapsed *fullNode, cached *fullNode) {
	// Hash the full node's children, caching the newly hashed subtrees
	cached = n.copy()
	collapsed = n.copy()
	for i := 0; i < 16; i++ {
		if child := n.Children[i]; child != nil {
			collapsed.Children[i], cached.Children[i] = h.hash(child, false)
		} else {
			collapsed.Children[i] = nilValueNode
		}
	}
	cached.Children[16] = n.Children[16]
	return collapsed, cached
}

// shortnodeToHash creates a hashNode from a shortNode. The supplied shortnode
// should have hex-type Key, which will be converted (without modification)
// into compact form for RLP encoding.
// If the rlp data is smaller than 32 bytes, `nil` is returned.
func (h *pureHasher) shortnodeToHash(n *shortNode, force bool) node {
	h.tmp.Reset()
	if err := rlp.Encode(&h.tmp, n); err != nil {
		panic("encode error: " + err.Error())
	}

	if len(h.tmp) < 32 && !force {
		return n // Nodes smaller than 32 bytes are stored inside their parent
	}
	return h.hashData(h.tmp)
}

// shortnodeToHash is used to creates a hashNode from a set of hashNodes, (which
// may contain nil values)
func (h *pureHasher) fullnodeToHash(n *fullNode, force bool) node {
	h.tmp.Reset()
	// Generate the RLP encoding of the node
	if err := n.EncodeRLP(&h.tmp); err != nil {
		panic("encode error: " + err.Error())
	}

	if len(h.tmp) < 32 && !force {
		return n // Nodes smaller than 32 bytes are stored inside their parent
	}
	return h.hashData(h.tmp)
}

// hashData hashes the provided data
func (h *pureHasher) hashData(data []byte) hashNode {
	n := make(hashNode, 32)
	h.sha.Reset()
	h.sha.Write(data)
	h.sha.Read(n)
	return n
}

// Code generated by github.com/fjl/gencodec. DO NOT EDIT.

package main

import (
	"encoding/json"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/common/math"
)

var _ = (*beaconBlockMarshaling)(nil)

// MarshalJSON marshals as JSON.
func (b beaconBlock) MarshalJSON() ([]byte, error) {
	type beaconBlock struct {
		ParentHash    common.Hash           `json:"parent_hash"    gencodec:"required"`
		FeeRecipient  common.Address        `json:"fee_recipient"  gencodec:"required"`
		StateRoot     common.Hash           `json:"state_root"     gencodec:"required"`
		ReceiptsRoot  common.Hash           `json:"receipts_root"  gencodec:"required"`
		LogsBloom     hexutil.Bytes         `json:"logs_bloom"     gencodec:"required"`
		Random        common.Hash           `json:"prev_randao"    gencodec:"required"`
		Number        math.HexOrDecimal64   `json:"block_number"   gencodec:"required"`
		GasLimit      math.HexOrDecimal64   `json:"gas_limit"      gencodec:"required"`
		GasUsed       math.HexOrDecimal64   `json:"gas_used"       gencodec:"required"`
		Timestamp     math.HexOrDecimal64   `json:"timestamp"     gencodec:"required"`
		ExtraData     hexutil.Bytes         `json:"extra_data"     gencodec:"required"`
		BaseFeePerGas *math.HexOrDecimal256 `json:"base_fee_per_gas" gencodec:"required"`
		BlockHash     common.Hash           `json:"block_hash"     gencodec:"required"`
		Transactions  []hexutil.Bytes       `json:"transactions"  gencodec:"required"`
	}
	var enc beaconBlock
	enc.ParentHash = b.ParentHash
	enc.FeeRecipient = b.FeeRecipient
	enc.StateRoot = b.StateRoot
	enc.ReceiptsRoot = b.ReceiptsRoot
	enc.LogsBloom = b.LogsBloom
	enc.Random = b.Random
	enc.Number = math.HexOrDecimal64(b.Number)
	enc.GasLimit = math.HexOrDecimal64(b.GasLimit)
	enc.GasUsed = math.HexOrDecimal64(b.GasUsed)
	enc.Timestamp = math.HexOrDecimal64(b.Timestamp)
	enc.ExtraData = b.ExtraData
	enc.BaseFeePerGas = (*math.HexOrDecimal256)(b.BaseFeePerGas)
	enc.BlockHash = b.BlockHash
	if b.Transactions != nil {
		enc.Transactions = make([]hexutil.Bytes, len(b.Transactions))
		for k, v := range b.Transactions {
			enc.Transactions[k] = v
		}
	}
	return json.Marshal(&enc)
}

// UnmarshalJSON unmarshals from JSON.
func (b *beaconBlock) UnmarshalJSON(input []byte) error {
	type beaconBlock struct {
		ParentHash    *common.Hash          `json:"parent_hash"    gencodec:"required"`
		FeeRecipient  *common.Address       `json:"fee_recipient"  gencodec:"required"`
		StateRoot     *common.Hash          `json:"state_root"     gencodec:"required"`
		ReceiptsRoot  *common.Hash          `json:"receipts_root"  gencodec:"required"`
		LogsBloom     *hexutil.Bytes        `json:"logs_bloom"     gencodec:"required"`
		Random        *common.Hash          `json:"prev_randao"    gencodec:"required"`
		Number        *math.HexOrDecimal64  `json:"block_number"   gencodec:"required"`
		GasLimit      *math.HexOrDecimal64  `json:"gas_limit"      gencodec:"required"`
		GasUsed       *math.HexOrDecimal64  `json:"gas_used"       gencodec:"required"`
		Timestamp     *math.HexOrDecimal64  `json:"timestamp"     gencodec:"required"`
		ExtraData     *hexutil.Bytes        `json:"extra_data"     gencodec:"required"`
		BaseFeePerGas *math.HexOrDecimal256 `json:"base_fee_per_gas" gencodec:"required"`
		BlockHash     *common.Hash          `json:"block_hash"     gencodec:"required"`
		Transactions  []hexutil.Bytes       `json:"transactions"  gencodec:"required"`
	}
	var dec beaconBlock
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}
	if dec.ParentHash == nil {
		return errors.New("missing required field 'parent_hash' for beaconBlock")
	}
	b.ParentHash = *dec.ParentHash
	if dec.FeeRecipient == nil {
		return errors.New("missing required field 'fee_recipient' for beaconBlock")
	}
	b.FeeRecipient = *dec.FeeRecipient
	if dec.StateRoot == nil {
		return errors.New("missing required field 'state_root' for beaconBlock")
	}
	b.StateRoot = *dec.StateRoot
	if dec.ReceiptsRoot == nil {
		return errors.New("missing required field 'receipts_root' for beaconBlock")
	}
	b.ReceiptsRoot = *dec.ReceiptsRoot
	if dec.LogsBloom == nil {
		return errors.New("missing required field 'logs_bloom' for beaconBlock")
	}
	b.LogsBloom = *dec.LogsBloom
	if dec.Random == nil {
		return errors.New("missing required field 'prev_randao' for beaconBlock")
	}
	b.Random = *dec.Random
	if dec.Number == nil {
		return errors.New("missing required field 'block_number' for beaconBlock")
	}
	b.Number = uint64(*dec.Number)
	if dec.GasLimit == nil {
		return errors.New("missing required field 'gas_limit' for beaconBlock")
	}
	b.GasLimit = uint64(*dec.GasLimit)
	if dec.GasUsed == nil {
		return errors.New("missing required field 'gas_used' for beaconBlock")
	}
	b.GasUsed = uint64(*dec.GasUsed)
	if dec.Timestamp == nil {
		return errors.New("missing required field 'timestamp' for beaconBlock")
	}
	b.Timestamp = uint64(*dec.Timestamp)
	if dec.ExtraData == nil {
		return errors.New("missing required field 'extra_data' for beaconBlock")
	}
	b.ExtraData = *dec.ExtraData
	if dec.BaseFeePerGas == nil {
		return errors.New("missing required field 'base_fee_per_gas' for beaconBlock")
	}
	b.BaseFeePerGas = (*big.Int)(dec.BaseFeePerGas)
	if dec.BlockHash == nil {
		return errors.New("missing required field 'block_hash' for beaconBlock")
	}
	b.BlockHash = *dec.BlockHash
	if dec.Transactions == nil {
		return errors.New("missing required field 'transactions' for beaconBlock")
	}
	b.Transactions = make([][]byte, len(dec.Transactions))
	for k, v := range dec.Transactions {
		b.Transactions[k] = v
	}
	return nil
}
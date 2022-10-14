package math

import (
	"math/big"
	"testing"
)

// TestFastModexp tests some cases found during fuzzing.
func TestFastModexp(t *testing.T) {
	for i, tc := range []struct {
		base string
		exp  string
		mod  string
	}{
		{"0xeffffff900002f00", "0x40000000000000", "0x200"},
		{"0xf000", "0x4f900b400080000", "0x400000d9d9d9d9d9d9d9d9d9d9d9d9d9d9d9d9d9d9d9d9d9d9d9d9d9d9d9d9ffffff005aeffd310000000000000000000000000000000000009f9f9f9f0000000000000000000000000800000000000000000000000000000000000000000000000000000000000000080000000000000000000000000000000000000000000000000000000000000befffa5a5a5fff900002f000040000000000000000000000000000000029d9d9d000000000000009f9f9f00000000000000009f9f9f000000f3a080ab00000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f0000000000002900009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f000000cf000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f000000000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f000000000000000000000000000000000000000000000000000000000000000800000000000000000000000000000000000000000000000000000000000000080000000000000000000000000000000000000000000000000000000000000befffff0000c0800000000800000000000000000000000000000002000000000000009f9f9f0000000000000000008000ff000000000000000000000000000000000000000000000000000000000000000800000000000000000000000000000000000000000000000000000000000000080000000000000000000000000000000000000000000000000000000000000befffa5a5a5fff900002f000040000000000000000000000000000000029d9d9d000000000000009f9f9f00000000000000009f9f9f000000f3a080ab00000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f0000000000002900009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f000000000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f000000000000000000000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f000000000000000000000000000000000000000000000000000000000000000800000000000000000000000000000000000000000000000000000000000000080000000000000000000000000000000000000000000000000000000000000beffffff900002f0000400000c100000000000000000000000000000000000000006160600000000000000000008000ff0000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f00000000000000009f9f0000"},
		{"5", "1435700818", "72"},
	} {
		var (
			base, _ = new(big.Int).SetString(tc.base, 0)
			exp, _  = new(big.Int).SetString(tc.exp, 0)
			mod, _  = new(big.Int).SetString(tc.mod, 0)
		)
		var a = FastExp(new(big.Int).Set(base), new(big.Int).Set(exp), new(big.Int).Set(mod))
		var b = new(big.Int).Exp(base, exp, mod)
		if a.Cmp(b) != 0 {
			t.Errorf("test %d: %#x ^ %#x mod %#x \n have %#x\n want %#x", i, base, exp, mod, a, b)
		}
	}
}

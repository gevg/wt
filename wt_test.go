package wt

import "testing"

func TestNew(t *testing.T) {
	tests := []struct {
		data   []uint8
		levels int
		bv     []uint64
	}{
		{
			[]uint8{0, 1, 6, 7, 1, 5, 4, 2, 6, 3},
			3,
			[]uint64{0b0011011010_0001111001_0110110010_0000000000_0000000000_0000000000_0000},
		},
	}

	for _, tc := range tests {
		wt, err := New(tc.data, tc.levels)
		if err != nil {
			t.Fatalf("creation wt failed: %s\n", err)
		}
		t.Logf("wt = %b\n", wt.bv)
	}
}

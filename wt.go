package wt

import "math/bits"

// Tree is a wavelet tree.
type Tree struct {
	n      int      // number of data elements
	levels int      // number of levels of tree // Kan ik hier height van maken??? Is dat een betere naam?
	bv     []uint64 // bit vector representation of tree
}

// New returns a wavelet tree with depth 'levels', created from 'data'.
func New(data []uint8, levels int) (*Tree, error) { // Misschien de Valsorda truc toepassen!
	n := len(data)
	bv := make([]uint64, (levels*n+63)>>6) // bit vector representation of tree
	hist := make([]int, 1<<levels)         // histogram of data (re-used at each tree level)

	var sigma uint8
	for i, v := range data {
		sigma |= v
		hist[v]++
		bv[i>>6] |= bit(levels-1+0, v) << (63 - i&63)
	}
	levels = bits.Len8(sigma)
	spos := make([]int, 1<<levels) // start positions of intervals in bit vector. // Hoeft niet noodzakelijk een int te zijn! Hangt af van n.

	for l := levels - 1; l > 0; l-- {
		for i := 0; i < 1<<l; i++ {
			hist[i] = hist[2*i] + hist[2*i+1]
		}
		for i := 1; i < 1<<l; i++ {
			spos[i] = spos[i-1] + hist[i-1]
		}
		for _, v := range data {
			spos[prefix(levels-1+l, v)]++
			pos := spos[prefix(levels-1+l, v)]
			bv[l*n+pos] = bit(levels-1+l, v)
		}
	}
	return nil, nil
}

func bit(i int, value uint8) uint64 {
	return uint64((value >> i) & 1)
}

func prefix(i int, value uint8) uint8 {
	return value >> i
}

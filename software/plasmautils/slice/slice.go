package slice

import (
	"sort"

	"../primeset"
)

type Slice struct {
	Begin uint32
	End   uint32
}

const TREE_HEIGHT = 24

func (s *Slice) GetAlignedSlices() []uint32 {
	res := []uint32{}
	a := s.Begin
	b := s.End
	q := uint32(1)
	t := uint32(0)
	for a+q < b {
		if a>>t&1 == 1 {
			res = append(res, 1<<(TREE_HEIGHT-t)+a/q-1)
			a += q
		}
		q <<= 1
		t++
	}
	for t+1 > 0 {
		if a+q <= b {
			res = append(res, 1<<(TREE_HEIGHT-t)+a/q-1)
			a += q
		}
		q >>= 1
		t--
	}
	return res
}

func LogProofInclusion(n []uint32) []uint32 {
	res := make([]uint32, 0)

	for _, item := range n {
		palpha := primeset.PrimeN(int(item * 2))
		res = append(res, palpha)
		for beta := item; true; beta = (beta - 1) >> 1 {
			pbeta := primeset.PrimeN(int(beta*2 + 1))
			res = append(res, pbeta)
			if beta == 0 {
				break
			}
		}
	}

	sort.Slice(res, func(i, j int) bool { return res[i] < res[j] })
	j := 0
	for i := 0; i < len(res); i++ {
		if res[j] != res[i] {
			j++
			res[j] = res[i]
		}
	}

	return res[0:j]
}

func LogProofExclusion(n []uint32) []uint32 {
	res := make([]uint32, 0)
	for _, item := range n {
		pbeta := primeset.PrimeN(int(item*2 + 1))
		res = append(res, pbeta)
		for alpha := item; true; alpha = (alpha - 1) >> 1 {
			palpha := primeset.PrimeN(int(alpha * 2))
			res = append(res, palpha)
			if alpha == 0 {
				break
			}
		}
	}
	sort.Slice(res, func(i, j int) bool { return res[i] < res[j] })
	j := 0
	for i := 0; i < len(res); i++ {
		if res[j] != res[i] {
			j++
			res[j] = res[i]
		}
	}

	return res[0:j]
}

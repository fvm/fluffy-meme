package solution

import (
	"sort"
)

type Developer struct {
	FrontendValue int
	BackendValue  int
}

type ByFrontEndValue []Developer

func (b ByFrontEndValue) Len() int {
	return len(b)
}

// If the relative value is high, i.e. large front, low back
func (b ByFrontEndValue) Less(i, j int) bool {
	if b[i].FrontendValue-b[i].BackendValue < b[j].FrontendValue-b[j].BackendValue {
		return true
	}
	if b[i].FrontendValue-b[i].BackendValue == b[j].FrontendValue-b[j].BackendValue {
		if b[i].FrontendValue < b[j].FrontendValue {
			return true
		}
		if b[i].FrontendValue == b[j].FrontendValue {
			return b[i].BackendValue < b[j].BackendValue
		}

	}
	return false
}

func (b ByFrontEndValue) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func Solution(A []int, B []int, F int) int {
	var d []Developer
	for i := 0; i < len(A); i++ {
		d = append(d, Developer{
			FrontendValue: A[i],
			BackendValue:  B[i],
		})
	}
	// [0] is high
	sort.Sort(sort.Reverse(ByFrontEndValue(d)))

	f, b := d[:F], d[F:]
	fv, _ := sum(f)
	_, bv := sum(b)
	return fv + bv
}

func sum(devs []Developer) (int, int) {
	var fv, bv int
	for _, dev := range devs {
		fv += dev.FrontendValue
		bv += dev.BackendValue
	}
	return fv, bv
}

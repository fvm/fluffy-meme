package solution

import "container/ring"

func Solution(A []int, K int) []int {
	if len(A) == 0 || K == 0 {
		return A
	}

	// write your code in Go 1.4
	r := ring.New(len(A))
	for i := range A {
		r.Value = A[i]
		r = r.Next()
	}
	r = r.Move(K)
	for i := 0; i < r.Len(); i++ {
		A[i] = r.Value.(int)
		r = r.Next()
	}
	return A
}

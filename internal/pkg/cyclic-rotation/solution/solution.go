package solution

import "container/ring"

func Solution(A []int, K int) []int {
	if len(A) == 0 || K == 0 {
		return A
	}

	r := ring.New(len(A))
	for i := range A {
		r.Value = A[i]
		r = r.Next()
	}
	r = r.Move(-1 * K) // Move **RIGHT** (not left)
	for i := 0; i < r.Len(); i++ {
		A[i] = r.Value.(int)
		r = r.Next()
	}
	return A
}

package solution

import "container/ring"

func Solution(A []int, K int) []int {
	// Boooooo!
	if len(A) == 0 || K == 0 {
		return A
	}

	r := ring.New(len(A))
	// Add all the ints
	for i := range A {
		r.Value = A[i]
		r = r.Next()
	}

	// Move **RIGHT** (not left)
	r = r.Move(-1 * K)

	// Remove all the ints
	for i := 0; i < r.Len(); i++ {
		A[i] = r.Value.(int)
		r = r.Next()
	}

	return A
}

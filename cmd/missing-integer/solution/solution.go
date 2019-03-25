package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

import "sort"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

//Solution Solves your problem of finding the lowest non-occurring positive integer
func Solution(A []int) int {
	var max = 0
	if len(A) == 0 { // Yeah, we're done
		return max + 1
	}

	// Speed should depend on the speed of sort and search

	// Right, so the easiest would be to first sort
	sort.Ints(A)

	// Now find the first positive number
	i := sort.Search(len(A), func(i int) bool { return A[i] > 0 })
	if i == len(A) { // Everything is smaller than 0, nothing remained
		return max + 1
	}

	// Drop the first part
	A = A[i:]

	for j := range A {

		// If more than 1 higher, it's discontinuous
		if A[j] > max+1 {
			break
		}
		// Leave it for what it is (or 1 higher). We can have dupes.
		max = A[j]
	}
	// We either reached the end or exited the loop
	return max + 1
}

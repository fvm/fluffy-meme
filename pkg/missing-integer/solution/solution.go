package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

import "sort"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

//Solution Solves your problem of finding the lowest non-occurring positive integer. Version two
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

/*
package solution_test

import (
	"github.com/fvm/fluffy-meme/pkg/missing-integer/solution"
	"testing"
)

func TestSolution(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "one", args: args{A: []int{1, 2, 3}}, want: 4},
		{name: "one", args: args{A: []int{1, 3, 6, 4, 1, 2}}, want: 5},
		{name: "one", args: args{A: []int{-1, -3}}, want: 1},
		{name: "one", args: args{A: []int{2}}, want: 1},
		{name: "one", args: args{A: []int{-1000, 1000}}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solution.Solution(tt.args.A); got != tt.want {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
*/

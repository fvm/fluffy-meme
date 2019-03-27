package solution_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/fvm/fluffy-meme/internal/pkg/lessons/cyclic-rotation/solution"
)

func TestSolution(t *testing.T) {
	type args struct {
		A []int
		K int
	}
	type test struct {
		name string
		args args
		want []int
	}

	gen := func(A []int, K int, want []int) test {
		return test{name: fmt.Sprintf("Test (%v, %d) => %v", A, K, want), args: args{A: A, K: K}, want: want}
	}

	tests := []test{
		gen([]int{3, 8, 9, 7, 6}, 3, []int{9, 7, 6, 3, 8}),
		gen([]int{0, 0, 0}, 1, []int{0, 0, 0}),
		gen([]int{1, 2, 3, 4}, 4, []int{1, 2, 3, 4}),
		gen([]int{1, 2, 3}, 3, []int{1, 2, 3}),
		gen([]int{1, 2}, 2, []int{1, 2}),
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solution.Solution(tt.args.A, tt.args.K); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}

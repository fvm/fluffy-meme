package solution_test

import (
	"fmt"
	"testing"

	"github.com/fvm/fluffy-meme/internal/pkg/odd-occurences-in-array/solution"
)

func TestSolution(t *testing.T) {
	type args struct {
		A []int
	}
	type test struct {
		name string
		args args
		want int
	}
	gen := func(A []int, i int) test {
		return test{
			name: fmt.Sprintf("poep-%d", i),
			args: args{
				A: A,
			},
			want: i,
		}
	}
	tests := []test{
		gen([]int{9, 3, 9, 3, 9, 7, 9}, 7),
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solution.Solution(tt.args.A); got != tt.want {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}

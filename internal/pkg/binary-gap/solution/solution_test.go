package solution_test

import (
	"fmt"
	"testing"

	"github.com/fvm/fluffy-meme/internal/pkg/binary-gap/solution"
)

func TestSolution(t *testing.T) {
	type args struct {
		N int
	}
	type test struct {
		name string
		args args
		want int
	}
	gen := func(i int, ans int) test {
		return test{
			name: fmt.Sprintf("%d => %d", i, ans),
			args: args{N: i},
			want: ans,
		}
	}
	tests := []test{
		gen(9, 2),
		gen(529, 4),
		gen(20, 1),
		gen(15, 0),
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solution.Solution(tt.args.N); got != tt.want {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}

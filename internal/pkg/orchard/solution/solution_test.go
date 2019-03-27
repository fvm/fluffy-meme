package solution_test

import (
	"testing"

	"github.com/fvm/fluffy-meme/internal/pkg/orchard/solution"
)

func TestSolution(t *testing.T) {
	type args struct {
		A []int
		K int
		L int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "Impossible", args: args{A: []int{}, K: 2, L: 2}, want: -1},
		{name: "Impossible", args: args{A: []int{10, 9, 15}, K: 2, L: 2}, want: -1},
		{name: "Exact", args: args{A: []int{10, 9, 15, 16}, K: 2, L: 2}, want: 50},
		{name: "Exact", args: args{A: []int{10, 9, 15, 16}, K: 1, L: 3}, want: 50},
		{name: "Regular", args: args{A: []int{10, 9, 1, 15, 16}, K: 1, L: 2}, want: 41},
		{name: "Regular", args: args{A: []int{10, 9, 1, 15, 16, 1, 1, 1, 1, 1, 32}, K: 1, L: 2}, want: 63},
		{name: "Regular", args: args{A: []int{10, 9, 1, 15, 1, 1, 16, 1, 1, 5, 1, 1, 32}, K: 5, L: 4}, want: 75},
		{name: "Regular", args: args{A: []int{10, 9, 1, 1, 1, 1, 16, 1, 1, 5, 1, 1, 32, 1, 1, 1, 1, 1, 5, 5, 5, 5, 5, 5, 5, 5, 40}, K: 5, L: 4}, want: 99},
		{name: "Regular", args: args{A: []int{10, 9, 1, 1, 1, 1, 16, 1, 1, 5, 1, 1, 32, 1, 1, 1, 1, 1, 5, 5, 5, 5, 5, 5, 5, 5, 40}, K: 5, L: 10}, want: 121},
		{name: "Regular", args: args{A: []int{10, 9, 1, 1, 1, 1, 16, 1, 1, 5, 1, 1, 32, 1, 1, 1, 1, 1, 5, 5, 5, 5, 5, 5, 5, 5, 1}, K: 12, L: 10}, want: 113},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solution.Solution(tt.args.A, tt.args.K, tt.args.L); got != tt.want {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}

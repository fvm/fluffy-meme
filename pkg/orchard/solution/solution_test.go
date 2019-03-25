package solution_test

import (
	"testing"

	"github.com/fvm/fluffy-meme/pkg/orchard/solution"
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
		{name: "Impossiple", args: args{A: []int{10, 9, 15}, K: 2, L: 2}, want: -1},
		{name: "Exact", args: args{A: []int{10, 9, 15, 16}, K: 2, L: 2}, want: 50},
		{name: "Exact", args: args{A: []int{10, 9, 15, 16}, K: 1, L: 3}, want: 50},
		{name: "Actual", args: args{A: []int{10, 9, 1, 15, 16}, K: 1, L: 2}, want: 41},
		{name: "Actual", args: args{A: []int{10, 9, 1, 15, 16, 1, 1, 1, 1, 1, 32}, K: 1, L: 2}, want: 63},
		{name: "Actual", args: args{A: []int{10, 9, 1, 15, 16, 1, 1, 1, 1, 1, 32}, K: 3, L: 2}, want: 65},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solution.Solution(tt.args.A, tt.args.K, tt.args.L); got != tt.want {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}

package solution_test

import (
	"github.com/fvm/fluffy-meme/cmd/MissingInteger/solution"
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

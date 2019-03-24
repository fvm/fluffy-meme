package solution_test

import (
	"github.com/fvm/fluffy-meme/pkg/fish/solution"
	"testing"
)

func TestSolution(t *testing.T) {
	type args struct {
		A []int
		B []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "test", args: args{A: []int{4, 3, 2, 1, 5}, B: []int{0, 1, 0, 0, 0}}, want: 2},
		{name: "test", args: args{A: []int{4, 3}, B: []int{0, 1}}, want: 2},
		{name: "test", args: args{A: []int{4, 3}, B: []int{1, 0}}, want: 1},
		{name: "test", args: args{A: []int{4}, B: []int{0}}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solution.Solution(tt.args.A, tt.args.B); got != tt.want {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkSolution(b *testing.B) {
	type args struct {
		A []int
		B []int
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "bench", args: args{
			A: []int{4, 3, 2, 1, 5, 3, 2, 1, 5, 3, 2, 1, 5, 3, 2, 1, 5, 3, 2, 1, 5, 3, 2, 1, 5, 3, 2, 1, 5, 3, 2, 1, 5, 3, 2, 1, 5, 3, 2, 1, 5, 3, 2, 1, 5, 3, 2, 1, 5},
			B: []int{0, 1, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0}}},
	}
	for _, bb := range tests {
		for n := 0; n < b.N; n++ {
			b.Run(bb.name, func(b *testing.B) {
				_ = solution.Solution(bb.args.A, bb.args.B)
			})
		}
	}
}

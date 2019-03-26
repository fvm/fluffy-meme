package solution_test

import (
	"github.com/fvm/fluffy-meme/pkg/winter-summer/solution"
	"testing"
)

func TestSolution(t *testing.T) {
	type args struct {
		T []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: string(""), args: args{T: []int{5, -2, 3, 8, 6}}, want: 3},
		{name: string(""), args: args{T: []int{-5, -5, -5, -42, 6, 12}}, want: 4},
		{name: string(""), args: args{T: []int{1, 1}}, want: 1},
		{name: string(""), args: args{T: []int{2, 1, 1}}, want: 2},
		{name: string(""), args: args{T: []int{-1, 1}}, want: 1},
		{name: string(""), args: args{T: []int{1000, 2, 1, -10000, 1, 10, 1001}}, want: 6},
		{name: string(""), args: args{T: []int{-1, 1}}, want: 1},
		{name: string(""), args: args{T: []int{-1, 1}}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solution.Solution(tt.args.T); got != tt.want {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}

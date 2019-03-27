package solution_test

import (
	"testing"

	"github.com/fvm/fluffy-meme/internal/pkg/f2b2f/solution"
)

func TestSolution(t *testing.T) {
	type args struct {
		A int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "", args: args{0}, want: 0},
		{name: "", args: args{1}, want: 1},
		{name: "", args: args{100000001}, want: 110000000},
		{name: "", args: args{12}, want: 12},
		{name: "", args: args{121}, want: 112},
		{name: "", args: args{112}, want: 121},
		{name: "", args: args{1123}, want: 1312},
		{name: "", args: args{123456}, want: 162534},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solution.Solution(tt.args.A); got != tt.want {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
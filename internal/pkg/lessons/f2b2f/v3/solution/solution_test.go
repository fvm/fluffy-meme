package solution_test

import (
	"testing"

	"github.com/fvm/fluffy-meme/internal/pkg/lessons/f2b2f/v3/solution"
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
		{name: "", args: args{12}, want: 12},
		{name: "", args: args{123}, want: 132},
		{name: "", args: args{1234}, want: 1423},
		{name: "", args: args{12345}, want: 15243},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solution.Solution(tt.args.A); got != tt.want {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}

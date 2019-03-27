package solution_test

import (
	"fmt"
	"testing"

	"github.com/fvm/fluffy-meme/internal/pkg/lessons/chocolates-by-numbers/solution"
)

func TestSolution(t *testing.T) {
	type args struct {
		N int
		M int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: " 10 , 3 =>  10", args: args{N: 10, M: 3}, want: 10},
		{name: " 11 , 3 =>  11", args: args{N: 11, M: 3}, want: 11},
		{name: " 10 , 2 =>   5", args: args{N: 10, M: 2}, want: 5},
		{name: " 10 , 4 =>   5", args: args{N: 10, M: 4}, want: 5},
		{name: "999 , 3 => 333", args: args{N: 999, M: 3}, want: 333},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solution.Solution(tt.args.N, tt.args.M); got != tt.want {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkSolution(b *testing.B) {
	type args struct {
		N int
		M int
	}
	type bench struct {
		name string
		args args
	}

	genBench := func(i int) bench {
		return bench{
			name: fmt.Sprintf("Bench-%d", i), args: args{
				N: 0,
				M: 0,
			},
		}
	}
	generateBenches := func(i []int) []bench {
		var benches = make([]bench, len(i))
		for j, k := range i {
			benches[j] = genBench(k)
		}
		return benches
	}
	var benches = generateBenches([]int{})
	for _, bb := range benches {
		b.Run(bb.name, func(b *testing.B) {
			_ = solution.Solution(bb.args.N, bb.args.M)
		})
	}
}

package solution_test

import (
	"fmt"
	"github.com/fvm/fluffy-meme/pkg/chocolates-by-numbers/solution"
	"testing"
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
		{name: "poep", args: args{N: 10, M: 4}, want: 5},
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

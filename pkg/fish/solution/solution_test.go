package solution_test

import (
	"fmt"
	"github.com/fvm/fluffy-meme/pkg/fish/solution"
	"math/rand"
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
		{name: "test", args: args{A: []int{4, 3, 2, 1, 5}, B: []int{0, 1, 0, 1, 0}}, want: 2},
		{name: "test", args: args{A: []int{4, 3, 2, 1, 5}, B: []int{0, 1, 1, 1, 0}}, want: 2},
		{name: "test", args: args{A: []int{4, 3, 2, 5, 1}, B: []int{0, 0, 1, 1, 0}}, want: 4},
		{name: "test", args: args{A: []int{4, 3, 2, 5, 1}, B: []int{0, 0, 0, 0, 0}}, want: 5},
		{name: "test", args: args{A: []int{4, 3, 2, 5, 1}, B: []int{1, 1, 1, 1, 1}}, want: 5},
		{name: "test", args: args{A: []int{4, 3, 2, 1, 5}, B: []int{1, 0, 0, 0, 0}}, want: 1},
		{name: "test", args: args{A: []int{4, 3, 2, 1, 5}, B: []int{0, 0, 0, 0, 1}}, want: 5},
		{name: "test", args: args{A: []int{4, 3, 2, 1, 5}, B: []int{0, 0, 0, 1, 0}}, want: 4},
		{name: "test", args: args{A: []int{4, 3, 2, 1, 5, 6, 7, 8, 9, 10}, B: []int{0, 0, 0, 1, 0, 0, 0, 1, 0, 0}}, want: 8},
		{name: "test", args: args{A: []int{4, 3, 2, 1, 5, 6, 7, 8, 9, 10}, B: []int{0, 0, 0, 1, 0, 0, 0, 0, 0, 1}}, want: 9},
		{name: "test", args: args{A: []int{4, 3, 2, 1, 5, 6, 7, 10, 9, 8}, B: []int{0, 0, 0, 1, 0, 0, 0, 1, 0, 0}}, want: 7},
		{name: "test", args: args{A: []int{2, 3, 4, 1, 5, 6, 7, 10, 9, 8}, B: []int{0, 1, 0, 1, 0, 0, 0, 1, 0, 0}}, want: 6},
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
	type bench struct {
		name string
		args args
	}

	gen := func(i int) args {
		a := rand.Perm(i)
		b := make([]int, len(a))
		for j := range b {
			b[j] = a[j] % 3
		}
		return args{a, b}
	}

	genbenches := func(is []int) []bench {
		var benches []bench
		for _, i := range is {
			benches = append(benches, bench{
				name: fmt.Sprintf("bench-%d", i),
				args: gen(i),
			})
		}
		return benches
	}

	var benches = genbenches([]int{10, 100, 1000, 10000, 100000})

	for _, bb := range benches {
		b.Run(bb.name, func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				_ = solution.Solution(bb.args.A, bb.args.B)
			}
		})
	}
}

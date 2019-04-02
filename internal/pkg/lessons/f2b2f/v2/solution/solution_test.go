package solution_test

import (
	"fmt"
	"math"
	"math/rand"
	"testing"

	"github.com/fvm/fluffy-meme/internal/pkg/lessons/f2b2f/v2/solution"
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
		{name: "", args: args{-123456}, want: -162534},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solution.Solution(tt.args.A); got != tt.want {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkSolution(b *testing.B) {
	type args struct {
		A int
	}
	type bench struct {
		name string
		args args
	}

	genArgs := func(i int) args {
		a := rand.Intn(i)
		return args{a}
	}

	genBenches := func(is []int) []bench {
		var benches []bench
		for _, i := range is {
			benches = append(benches, bench{
				name: fmt.Sprintf("bench-%d", i),
				args: genArgs(i),
			})
		}
		return benches
	}

	r := rand.New(rand.NewSource(10101))

	var benchInput []int
	for i := 1; i <= 18; i++ {
		benchInput = append(benchInput, int(r.Float64()*math.Pow10(i)))
	}
	var benches = genBenches(benchInput)

	for _, bb := range benches {
		b.Run(bb.name, func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				_ = solution.Solution(bb.args.A)
			}
		})
	}
}

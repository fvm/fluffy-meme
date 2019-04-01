package f2b2f_test

import (
	"fmt"
	"math"
	"math/rand"
	"testing"

	v1 "github.com/fvm/fluffy-meme/internal/pkg/lessons/f2b2f/solution"
	v2 "github.com/fvm/fluffy-meme/internal/pkg/lessons/f2b2f/v2/solution"
	v3 "github.com/fvm/fluffy-meme/internal/pkg/lessons/f2b2f/v3/solution"
)

func BenchmarkSolution(b *testing.B) {
	type args struct {
		A int
	}
	type bench struct {
		name string
		args args
	}

	gen := func(i int) args {
		a := rand.Intn(i)
		return args{a}
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

	r := rand.New(rand.NewSource(10101))

	var benchInput []int
	for i := 1; i <= 12; i++ {
		for j := 0; j < 10; j++ {
			benchInput = append(benchInput, int(r.Float64()*math.Pow10(i)))
		}
	}
	var benches = genbenches(benchInput)

	for _, bb := range benches {
		b.Run(fmt.Sprintf("v1-%s", bb.name), func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				_ = v1.Solution(bb.args.A)
			}
		})
	}
	for _, bb := range benches {
		b.Run(fmt.Sprintf("v2-%s", bb.name), func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				_ = v2.Solution(bb.args.A)
			}
		})
	}
	for _, bb := range benches {
		b.Run(fmt.Sprintf("v3-%s", bb.name), func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				_ = v3.Solution(bb.args.A)
			}
		})
	}
}

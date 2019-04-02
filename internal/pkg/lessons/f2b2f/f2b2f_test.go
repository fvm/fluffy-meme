package f2b2f_test

import (
	"fmt"
	"math"
	"math/rand"
	"testing"

	s "github.com/fvm/fluffy-meme/internal/pkg/lessons/f2b2f/v3/solution"
)

//  In order to compare versions without changing the actual code, change the imported package.
//	This allows using benchcmp (golang.org/x/tools/cmd/benchcmp) to compare the benches

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
				_ = s.Solution(bb.args.A)
			}
		})
	}
}

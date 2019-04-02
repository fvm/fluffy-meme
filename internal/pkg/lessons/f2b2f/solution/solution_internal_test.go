package solution

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
	"testing"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

func randStringBytesMaskImpr(n int) []string {
	b := make([]byte, n)
	// A rand.Int63() generates 63 random bits, enough for letterIdxMax letters!
	rs := rand.NewSource(10101)
	for i, cache, remain := n-1, rs.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = rand.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return strings.Split(string(b), "")
}

func Benchmark_hustle(b *testing.B) {
	type args struct {
		S []string
	}
	type bench struct {
		name string
		args args
	}

	genBenches := func(is []int) []bench {
		var benches []bench
		for _, i := range is {
			args := args{S: randStringBytesMaskImpr(i)}
			benches = append(benches, bench{
				name: fmt.Sprintf("bench-%d", i),
				args: args,
			})
		}
		return benches
	}

	var benchInput []int
	for i := 1.0; i <= 16.0; i++ {
		benchInput = append(benchInput, int(math.Pow(2.0, i)))
	}

	var benches = genBenches(benchInput)

	for _, bb := range benches {
		b.Run(bb.name, func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				_ = hustle(bb.args.S)
			}
		})
	}
}

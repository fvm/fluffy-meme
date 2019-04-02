package solution

import (
	"fmt"
	"math"
	"math/rand"
	"testing"
)

const seed = 10101

// Copied from StackOverflow. Just copy-paste it everywhere
func randString(b *testing.B, length int, seed int) []string {
	b.Helper()

	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	const (
		letterIdxBits = 6                    // 6 bits to represent a letter index
		letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
		letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
	)

	bs := make([]string, length)
	rs := rand.NewSource(int64(seed))

	// A rand.Int63() generates 63 random bits, enough for letterIdxMax letters!
	for i, cache, remain := length-1, rs.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = rand.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			bs[i] = string(letterBytes[idx])
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return bs
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
			args := args{S: randString(b, i, seed)}
			benches = append(benches, bench{
				name: fmt.Sprintf("%d", i),
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

package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

type fish struct {
	idx       int
	size      int
	direction int
}

func Solution(A []int, B []int) int {
	// write your code in Go 1.4
	var river []fish

	for i := range A {
		river = append(river, fish{
			i,
			A[i],
			B[i],
		})
	}

	for t := 0; ; t++ { // time
		var remaining []fish
		for i := range remaining { //fishies
			if river[i].direction == 0 { // Upstream
				remaining = append(remaining, river[i])
				continue
			}

			if len(river)-i <= 1 { // Last fish
				remaining = append(remaining, river[i])
				continue
			}

			if river[i].direction == 1 { // Downstream
				if river[i+1].direction == 0 { // Collision
					if river[i].size < river[i+1].size {
						continue // drop this fishie
					}
				}

			}

		}
	}
	return 0
}

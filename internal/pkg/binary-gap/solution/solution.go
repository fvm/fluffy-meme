package solution

import (
	"fmt"
	"strings"
)

func Solution(N int) int {
	// write your code in Go 1.4
	s := fmt.Sprintf("%b", N)
	ss := strings.Split(s, "")
	current, max := 0, 0
	skip := true
	for i := len(s) - 1; i > 0; i-- {
		if ss[i] == "1" { // Any 1 will reset the count
			skip = false
			current = 0
			continue
		}
		if skip == true {
			continue
		}
		current++
		if current > max {
			max = current
		}
	}

	return max
}

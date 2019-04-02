package solution

import (
	"strconv"
	"strings"
)

func Solution(A int) int {

	number := strconv.Itoa(A)

	numbers := strings.Split(number, "")

	numbers = hustle(numbers)

	number = strings.Join(numbers, "")

	A, _ = strconv.Atoi(number) // <- Bad, but can't fix it

	return A
}

func hustle(original []string) []string {
	if len(original) <= 2 {
		return original
	}
	var shuffled = make([]string, len(original), len(original))
	for n, m := 0, len(original)-1; n <= m; n, m = n+1, m-1 {
		shuffled[2*n] = original[n]
		if n == m {
			break
		}
		shuffled[2*n+1] = original[m]

	}
	return shuffled
}

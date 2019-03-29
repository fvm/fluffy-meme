package solution

import (
	"strconv"
	"strings"
)

func Solution(A int) int {

	// Convert to string and split into 'chars'
	number := strconv.Itoa(A)
	numbers := strings.Split(number, "")
	if len(numbers) <= 2 {
		return A
	}
	var b []string
	for n, m := 0, len(numbers)-1; n <= m; n, m = n+1, m-1 {
		b = append(b, numbers[n])
		if n == m {
			break
		}
		b = append(b, numbers[m])
	}

	number = strings.Join(b, "")

	B, _ := strconv.Atoi(number) // I can't do anything with the error here

	return B
}

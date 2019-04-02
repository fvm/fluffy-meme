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
	numbers = hustle(numbers)

	number = strings.Join(numbers, "")

	A, _ = strconv.Atoi(number) // I can't do anything with the error here

	return A
}

func hustle(numbers []string) []string {
	var shuffled []string
	for n, m := 0, len(numbers)-1; n <= m; n, m = n+1, m-1 {
		shuffled = append(shuffled, numbers[n])
		if n == m {
			break
		}
		shuffled = append(shuffled, numbers[m])
	}
	return shuffled
}

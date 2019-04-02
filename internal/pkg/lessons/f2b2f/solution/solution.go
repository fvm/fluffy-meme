package solution

import (
	"strconv"
	"strings"
)

// Solution does a front to back to front
func Solution(A int) int {

	number := strconv.Itoa(A)

	numbers := strings.Split(number, "")

	numbers = hustle(numbers)

	number = strings.Join(numbers, "")

	A, _ = strconv.Atoi(number)

	return A

}

func hustle(numbers []string) []string {
	if len(numbers) <= 2 {
		return numbers
	}
	var shuffled []string
	for {
		head := numbers[0]
		shuffled = append(shuffled, head)
		if len(numbers) <= 1 {
			break
		}
		numbers = reverse(numbers[1:])
	}
	return shuffled
}

// Taken from the slice tricks wiki page
func reverse(s []string) []string {
	for left, right := 0, len(s)-1; left < right; left, right = left+1, right-1 {
		s[left], s[right] = s[right], s[left]
	}
	return s
}

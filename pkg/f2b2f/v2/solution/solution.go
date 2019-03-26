package solution

import (
	"strconv"
	"strings"
)

// Solution does a front to back to front
func Solution(A int) int {

	// Speed is not an issue, so we can go to char array. So no need for doing a base-10 split

	//var numbers []string
	number := strconv.Itoa(A)
	numbers := strings.Split(number, "")

	numbers = hustle(numbers)
	number = strings.Join(numbers, "")

	B, _ := strconv.Atoi(number) // I can't do anything with the error here

	return B
}

func hustle(s []string) []string {
	if len(s) <= 2 {
		return s
	}

	return append(s[:1], hustle(reverse(s[1:]))...)
}

// Taken from the slice tricks wiki page
func reverse(s []string) []string {
	for left, right := 0, len(s)-1; left < right; left, right = left+1, right-1 {
		s[left], s[right] = s[right], s[left]
	}

	return s
}

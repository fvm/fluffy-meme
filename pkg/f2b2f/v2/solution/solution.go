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
	// Prevent going out of bounds
	if len(s) == 0 {
		return s
	}
	// Catch a negative number by including the minus
	offset := 1
	if s[0] == "-" {
		offset++
	}
	// All combinations of 2 numbers or less will be returned in the same order
	if len(s) <= offset+1 {
		return s
	}
	// Otherwise return the first number followed by the first of the flipped rest (rinse and repeat)
	return append(s[:offset], hustle(reverse(s[offset:]))...)

}

// Taken from the slice tricks wiki page
func reverse(s []string) []string {
	for left, right := 0, len(s)-1; left < right; left, right = left+1, right-1 {
		s[left], s[right] = s[right], s[left]
	}

	return s
}

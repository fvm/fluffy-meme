package solution

import (
	"strconv"
	"strings"
)

// Solution solves the front-to-back-to-front "encryption" by recursively
// * Splitting the slice into a head and tail
// * Appending the result of calling itself on the flipped tail to the head
// * Until the tail is at most 2 elements long
func Solution(A int) int {

	// Convert to string and split into 'chars'
	number := strconv.Itoa(A)
	numbers := strings.Split(number, "")

	i := 0

	if numbers[0] == "-" { // skip the first
		i = 1
	}

	numbers = append(numbers[:i], hustle(numbers[i:])...)

	number = strings.Join(numbers, "")

	B, _ := strconv.Atoi(number) // I can't do anything with the error here

	return B
}

func hustle(s []string) []string {
	// All combinations of 2 numbers or less will be returned in the same order
	if len(s) <= 2 {
		return s
	}
	// Otherwise return the first number followed by the first of the flipped rest (rinse and repeat)
	return append(s[:1], hustle(reverse(s[1:]))...)

}

// Taken from the slice tricks wiki page
func reverse(s []string) []string {
	for left, right := 0, len(s)-1; left < right; left, right = left+1, right-1 {
		s[left], s[right] = s[right], s[left]
	}

	return s
}

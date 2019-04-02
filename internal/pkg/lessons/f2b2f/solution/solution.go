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

	var numbers = strings.Split(number, "")

	// "put my thing down, reverse it and reverse it"
	numbers = hustle(numbers)

	number = strings.Join(numbers, "")
	B, _ := strconv.Atoi(number) // I can't do anything with the error here
	return B

}

func hustle(numbers []string) []string {
	var shuffled []string
	for {
		// first one
		head := numbers[0]
		shuffled = append(shuffled, head)
		if len(numbers) <= 1 {
			// exit (that was the last one and we don't want to access numbers[1:]
			break
		}
		// Shift and reverse the tail
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

package solution

import "sort"

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(T []int) int {
	// write your code in Go 1.4

	// We can assume the input is at least 2 long (but let's be safe)
	if len(T) <= 2 {
		// It should at least be one day long
		return 1
	}

	// Alright, the array will start with winter
	// The first "day" will be the highest value in winter
	warmestDayOfWinter := T[0]

	// Sort all the ints, all lower values will be to the left
	sort.Ints(T)

	// Now find the first day warmer than that, all other values will be values of summer
	firstSummerDay := sort.Search(len(T), func(i int) bool {
		// Note: This uses binary search through a tree. Could get slow when the tree is unbalanced
		return T[i] > warmestDayOfWinter
	})

	// Edge case, we didn't find a summer
	if firstSummerDay >= len(T) {
		// But, summer must be at least one day long. The rest is winter
		return len(T) - 1
	}

	// Return the length of the first part
	winter := T[:firstSummerDay]
	return len(winter)
}

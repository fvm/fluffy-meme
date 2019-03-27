package solution

func Solution(A []int) int {
	hist := make(map[int]int)
	for _, i := range A {
		hist[i]++
	}
	for i := range hist {
		if hist[i]%2 == 1 {
			return i
		}
	}
	return 0
}

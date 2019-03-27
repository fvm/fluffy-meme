package solution

func Solution(N int, M int) int {
	if N%M == 0 {
		return N / M
	}

	// Get the greatest common divisor
	g := gcd(N, M)
	if g == 1 { // No gcd (primes, e.g. 10/3)
		return N
	}

	return Solution(N*g, M)
}
func gcd(a int, b int) int {
	if a == b {
		return a
	}
	if a > b {
		return gcd(a-b, b)
	}
	if a < b {
		return gcd(a, b-a)
	}
	return 1
}

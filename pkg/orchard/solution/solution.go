package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

type partition struct {
	value  int
	head   []int
	middle []int
	tail   []int
}

func Solution(A []int, K int, L int) int {
	// No disjoint set
	if len(A) < K+L {
		return -1
	}
	// Perfectly disjoint
	if len(A) == K+L { // Just split it
		return sum(A[:K]) + sum(A[K:])
	}

	// It's tempting to go for channels and goroutine, but that's complicated, so let's brute force it

	// Start grabbing pieces for Alice first
	piecesForAlice := divide(A, K)

	// Now start seeing if we can find something for Bob

	var pairs [][]partition

	for _, pieceForAlice := range piecesForAlice {
		var piecesForBob []partition
		if len(pieceForAlice.head) >= L {
			piecesForBob = append(piecesForBob, divide(pieceForAlice.head, L)...)
		}

		if len(pieceForAlice.tail) >= L {
			piecesForBob = append(piecesForBob, divide(pieceForAlice.tail, L)...)
		}
		// So now we have all pieces for Bob
		for _, pieceForBob := range piecesForBob {
			pairs = append(pairs, []partition{
				pieceForAlice,
				pieceForBob,
			})
		}
	}
	bestPair := pairs[0]
	mostApples := bestPair[0].value + bestPair[1].value
	for _, pair := range pairs {
		apples := pair[0].value + pair[1].value
		if apples >= mostApples {
			bestPair = pair
			mostApples = apples
		}
	}

	return mostApples
}

func sum(i []int) int {
	var s int
	for _, v := range i {
		s += v
	}
	return s
}

//divide Finds subsets of size 	K in A, returns a partition
func divide(A []int, K int) []partition {
	var partitions []partition
	if K > len(A) {
		return partitions
	}

	for min, max := 0, K; max <= len(A); min, max = min+1, max+1 {
		head, middle, tail := A[:min], A[min:max], A[max:]
		partitions = append(partitions, partition{
			sum(middle),
			head,
			middle,
			tail,
		})
	}
	return partitions
}

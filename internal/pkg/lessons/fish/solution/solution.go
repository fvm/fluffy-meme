package solution

import "log"

type fish struct {
	name      int
	size      int
	direction direction
}

type direction int

const (
	upstream direction = iota
	downstream
)

func Solution(A []int, B []int) int {

	// write your code in Go 1.4
	var fishies []fish
	if len(A) == 0 {
		return 0
	}

	for i := range A {
		fishies = append(fishies, fish{
			name:      i,
			size:      A[i],
			direction: direction(B[i]),
		})
	}

	s := stack{
		river:     []fish{},
		direction: downstream,
		stacks:    map[direction][]fish{upstream: {}, downstream: {}},
	}

	for {
		head := fishies[0]
		s.add(head)
		if len(fishies) <= 1 {
			s.compact()
			s.flush()
			break
		}
		fishies = fishies[1:]
	}
	return len(s.river)
}

func (s *stack) add(f fish) {
	if f.direction != s.direction {
		s.switchDirection()
	}
	s.stacks[s.direction] = append(s.stacks[s.direction], f)
}

type stack struct {
	river     []fish
	direction direction
	stacks    map[direction][]fish
}

func (s *stack) switchDirection() {
	switch s.direction {
	case downstream:
		s.direction = upstream
	case upstream:
		s.compact()
		s.direction = downstream
	}
}

func (s *stack) compact() {

	down, up := resolve(reverse(s.stacks[downstream]), s.stacks[upstream])

	// Either up or down will be empty
	if len(down) != 0 && len(up) != 0 {
		log.Fatal("Whups")
	}

	// Upstream go off into the river
	s.river = append(s.river, up...)
	// Downstream keep swimming
	s.stacks[upstream] = []fish{}
	s.stacks[downstream] = down

}

func (s *stack) flush() {
	s.river = append(s.river, s.stacks[upstream]...)
	s.river = append(s.river, s.stacks[downstream]...)
}

func reverse(f []fish) []fish {
	for left, right := 0, len(f)-1; left < right; left, right = left+1, right-1 {
		f[left], f[right] = f[right], f[left]
	}
	return f
}

func resolve(downstack []fish, upstack []fish) ([]fish, []fish) {
	if len(upstack) == 0 || len(downstack) == 0 {
		return downstack, upstack
	}

	for {
		downhead, uphead := downstack[0], upstack[0]
		if downhead.size < uphead.size {
			return resolve(downstack[1:], upstack)
		} else {
			return resolve(downstack, upstack[1:])
		}
	}

}

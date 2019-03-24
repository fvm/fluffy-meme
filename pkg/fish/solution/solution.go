package solution

type fish struct {
	name      int
	size      int
	direction stackselector
}

type mode int

type stackselector int

const (
	upstream stackselector = iota
	downstream
)

// mode can (later) be used to check for lock status
const (
	accumulate mode = iota
	compact
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
			direction: stackselector(B[i]),
		})
	}

	s := stack{
		river:         []fish{},
		stackselector: downstream,
		mode:          accumulate,
		stacks:        map[stackselector][]fish{upstream: {}, downstream: {}},
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
	switch f.direction {
	case upstream:
		if s.stackselector == downstream {
			s.switchStack()
		}
		s.stacks[upstream] = append(s.stacks[upstream], f)
	case downstream:
		if s.stackselector == upstream {
			s.switchStack()
		}
		s.stacks[downstream] = append(s.stacks[downstream], f)
	}
}

type stack struct {
	river         []fish
	stackselector stackselector
	mode          mode
	stacks        map[stackselector][]fish
}

func (s *stack) switchStack() {
	switch s.stackselector {
	case downstream:
		s.stackselector = upstream
	case upstream:
		defer func() {
			s.switchMode()
			s.compact()
			s.flush()
			s.switchMode()
		}()
		s.stackselector = downstream
	}
}

func (s *stack) switchMode() {
	switch s.mode {
	case accumulate:
		s.mode = compact
	case compact:
		s.mode = accumulate
	}
}

func (s *stack) compact() {
	s.river = append(s.river, resolve(s.stacks[upstream], s.stacks[downstream])...)
}

func (s *stack) flush() {
	s.stacks[upstream] = []fish{}
	s.stacks[downstream] = []fish{}
}

func resolve(ls []fish, rs []fish) []fish {
	if len(ls) == 0 {
		return rs
	}
	if len(rs) == 0 {
		return ls
	}

	for {
		lh, rh := ls[0], rs[0]
		if lh.size < rh.size {
			return resolve(ls[1:], rs)
		} else {
			return resolve(ls, rs[1:])
		}
	}

}

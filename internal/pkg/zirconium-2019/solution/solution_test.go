package solution_test

import (
	"fmt"
	"testing"

	s "github.com/fvm/fluffy-meme/internal/pkg/zirconium-2019/solution"
)

func TestSolution(t *testing.T) {
	type args struct {
		A []int
		B []int
		F int // Front-ends
	}
	type test struct {
		name string
		args args
		want int
	}
	gen := func(a args, w int) test {
		return test{name: fmt.Sprintf("%v", a), args: a, want: w}
	}
	tests := []test{
		gen(args{A: []int{4, 2, 1}, B: []int{2, 5, 3}, F: 2}, 10),
		gen(args{A: []int{7, 1, 4, 4}, B: []int{5, 3, 4, 3}, F: 2}, 18),
		gen(args{A: []int{5, 5, 5}, B: []int{5, 5, 5}, F: 1}, 15),
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := s.Solution(tt.args.A, tt.args.B, tt.args.F); got != tt.want {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_byFrontendValue_Less(t *testing.T) {
	type args struct {
		i int
		j int
	}

	tests := []struct {
		name string
		b    s.ByFrontEndValue
		args args
		want bool
	}{
		{
			name: "flups",
			b: s.ByFrontEndValue([]s.Developer{
				{
					BackendValue:  0,
					FrontendValue: 10,
				}, {
					BackendValue:  0,
					FrontendValue: 9,
				},
			},
			),
			args: args{
				i: 0, j: 1,
			},
			want: false,
		},
		{
			name: "flups",
			b: s.ByFrontEndValue([]s.Developer{
				{
					BackendValue:  0,
					FrontendValue: 10,
				}, {
					BackendValue:  9,
					FrontendValue: 9,
				},
			},
			),
			args: args{
				i: 0, j: 1,
			},
			want: false,
		},
		{
			name: "flups",
			b: s.ByFrontEndValue([]s.Developer{
				{
					BackendValue:  10,
					FrontendValue: 10,
				}, {
					BackendValue:  0,
					FrontendValue: 10,
				},
			},
			),
			args: args{
				i: 0, j: 1,
			},
			want: true,
		},
		{
			name: "flups",
			b: s.ByFrontEndValue([]s.Developer{
				{
					BackendValue:  0,
					FrontendValue: 1,
				}, {
					BackendValue:  0,
					FrontendValue: 1,
				},
			},
			),
			args: args{
				i: 0, j: 1,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.Less(tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("ByFrontendValue.Less() = %v, want %v", got, tt.want)
			}
		})
	}
}

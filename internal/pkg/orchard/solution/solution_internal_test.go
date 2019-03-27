package solution

import (
	"reflect"
	"testing"
)

func Test_divide(t *testing.T) {
	type args struct {
		A []int
		K int
	}
	tests := []struct {
		name string
		args args
		want []partition
	}{
		{
			name: "divide",
			args: args{A: []int{1}, K: 0},
			want: []partition{
				{value: 0, head: []int{}, middle: []int{}, tail: []int{1}},
				{value: 0, head: []int{1}, middle: []int{}, tail: []int{}},
			},
		},
		{
			name: "divide",
			args: args{A: []int{1}, K: 1},
			want: []partition{
				{value: 1, head: []int{}, middle: []int{1}, tail: []int{}},
			},
		},
		{
			name: "divide",
			args: args{A: []int{1, 2}, K: 1},
			want: []partition{
				{value: 1, head: []int{}, middle: []int{1}, tail: []int{2}},
				{value: 2, head: []int{1}, middle: []int{2}, tail: []int{}},
			},
		},
		{
			name: "divide",
			args: args{A: []int{1}, K: 2},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := divide(tt.args.A, tt.args.K); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("divide() = %v, want %v", got, tt.want)
			}
		})
	}
}

package rotate

import "testing"

func Equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func TestRotate(t *testing.T) {
	type args struct {
		s []int
		k uint
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "[], k = 0",
			args: args{[]int{}, 0},
			want: []int{},
		},
		{
			name: "[], k = 1",
			args: args{[]int{}, 1},
			want: []int{},
		},
		{
			name: "[1], k = 0",
			args: args{[]int{1}, 0},
			want: []int{1},
		},
		{
			name: "[1], k = 1",
			args: args{[]int{1}, 1},
			want: []int{1},
		},
		{
			name: "[1, 2, 3], k = 0",
			args: args{[]int{1, 2, 3}, 0},
			want: []int{1, 2, 3},
		},
		{
			name: "[1, 2, 3], k = 1",
			args: args{[]int{1, 2, 3}, 1},
			want: []int{3, 1, 2},
		},
		{
			name: "[1, 2, 3], k = 2",
			args: args{[]int{1, 2, 3}, 2},
			want: []int{2, 3, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := make([]int, len(tt.args.s))
			copy(got, tt.args.s)
			if rotate(got, tt.args.k); !Equal(got, tt.want) {
				t.Errorf("rotate(%v, %v) = %v, want %v", tt.args.s, tt.args.k, got, tt.want)
			}
		})
	}
}

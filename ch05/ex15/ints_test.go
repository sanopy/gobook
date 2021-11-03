package ints

import "testing"

func TestMinElement(t *testing.T) {
	type args struct {
		vals []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "{1}",
			args: args{[]int{1}},
			want: 1,
		},
		{
			name: "{1, 2}",
			args: args{[]int{1, 2}},
			want: 1,
		},
		{
			name: "{3, 5, 4, 1, 2}",
			args: args{[]int{3, 5, 4, 1, 2}},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := minElement(tt.args.vals...); got != tt.want {
				t.Errorf("minElement(%v) = %v, want %v", tt.args.vals, got, tt.want)
			}
		})
	}
}

func TestMinElementError(t *testing.T) {
	_, err := minElement()
	if err == nil {
		t.Errorf("must raise error")
	}
}

func TestMaxElement(t *testing.T) {
	type args struct {
		vals []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "{1}",
			args: args{[]int{1}},
			want: 1,
		},
		{
			name: "{1, 2}",
			args: args{[]int{1, 2}},
			want: 2,
		},
		{
			name: "{3, 5, 4, 1, 2}",
			args: args{[]int{3, 5, 4, 1, 2}},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := maxElement(tt.args.vals...); got != tt.want {
				t.Errorf("maxElement(%v) = %v, want %v", tt.args.vals, got, tt.want)
			}
		})
	}
}

func TestMaxElementError(t *testing.T) {
	_, err := maxElement()
	if err == nil {
		t.Errorf("must raise error")
	}
}

func TestMin(t *testing.T) {
	type args struct {
		vals []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "{1}",
			args: args{[]int{1}},
			want: 1,
		},
		{
			name: "{1, 2}",
			args: args{[]int{1, 2}},
			want: 1,
		},
		{
			name: "{3, 5, 4, 1, 2}",
			args: args{[]int{3, 5, 4, 1, 2}},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := min(tt.args.vals[0], tt.args.vals[1:]...); got != tt.want {
				t.Errorf("min(%v) = %v, want %v", tt.args.vals, got, tt.want)
			}
		})
	}
}

func TestMax(t *testing.T) {
	type args struct {
		vals []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "{1}",
			args: args{[]int{1}},
			want: 1,
		},
		{
			name: "{1, 2}",
			args: args{[]int{1, 2}},
			want: 2,
		},
		{
			name: "{3, 5, 4, 1, 2}",
			args: args{[]int{3, 5, 4, 1, 2}},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := max(tt.args.vals[0], tt.args.vals[1:]...); got != tt.want {
				t.Errorf("max(%v) = %v, want %v", tt.args.vals, got, tt.want)
			}
		})
	}
}

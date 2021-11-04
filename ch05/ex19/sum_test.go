package sum

import "testing"

func TestSum(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: `sum(0)`,
			args: args{0},
			want: 0,
		},
		{
			name: `sum(1)`,
			args: args{1},
			want: 1,
		},
		{
			name: `sum(10)`,
			args: args{10},
			want: 55,
		},
		{
			name: `sum(100)`,
			args: args{100},
			want: 5050,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sum(tt.args.x); got != tt.want {
				t.Errorf("sum(%v) = %v, want %v", tt.args.x, got, tt.want)
			}
		})
	}
}

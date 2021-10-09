package mycmplx

import (
	"testing"
)

func TestBigRatAdd(t *testing.T) {
	type args struct {
		a *BigRat
		b *BigRat
	}
	tests := []struct {
		name string
		args args
		want *BigRat
	}{
		{
			name: "1 + i",
			args: args{a: NewBigRat(1, 0), b: NewBigRat(0, 1)},
			want: NewBigRat(1, 1),
		},
		{
			name: "(1 + i) + (1 + i)",
			args: args{a: NewBigRat(1, 1), b: NewBigRat(1, 1)},
			want: NewBigRat(2, 2),
		},
		{
			name: "(3 + 4i) + (5 + 7i)",
			args: args{a: NewBigRat(3, 4), b: NewBigRat(5, 7)},
			want: NewBigRat(8, 11),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.a.Add(tt.args.b); !(got.Real.Cmp(tt.want.Real) == 0 && got.Imag.Cmp(tt.want.Imag) == 0) {
				t.Errorf("%v.Add(%v) = %v, want %v", tt.args.a, tt.args.b, got, tt.want)
			}
		})
	}
}

func TestBigRatMul(t *testing.T) {
	type args struct {
		a *BigRat
		b *BigRat
	}
	tests := []struct {
		name string
		args args
		want *BigRat
	}{
		{
			name: "1·i",
			args: args{a: NewBigRat(1, 0), b: NewBigRat(0, 1)},
			want: NewBigRat(0, 1),
		},
		{
			name: "(1 + i)(1 + i)",
			args: args{a: NewBigRat(1, 1), b: NewBigRat(1, 1)},
			want: NewBigRat(0, 2),
		},
		{
			name: "(2 + i)(-1 + 3i)",
			args: args{a: NewBigRat(2, 1), b: NewBigRat(-1, 3)},
			want: NewBigRat(-5, 5),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.a.Mul(tt.args.b); !(got.Real.Cmp(tt.want.Real) == 0 && got.Imag.Cmp(tt.want.Imag) == 0) {
				t.Errorf("%v.Mul(%v) = %v, want %v", tt.args.a, tt.args.b, got, tt.want)
			}
		})
	}
}

func TestBigRatAbs(t *testing.T) {
	type args struct {
		z *BigRat
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "|2|",
			args: args{z: NewBigRat(2, 0)},
			want: 2,
		},
		{
			name: "|2i|",
			args: args{z: NewBigRat(0, 2)},
			want: 2,
		},
		{
			name: "|3 + 4i|",
			args: args{z: NewBigRat(3, 4)},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.z.Abs(); got != tt.want {
				t.Errorf("%v.Abs() = %v, want %v", tt.args.z, got, tt.want)
			}
		})
	}
}

package mycmplx

import (
	"math/big"
	"testing"
)

func TestBigFloatAdd(t *testing.T) {
	type args struct {
		a *BigFloat
		b *BigFloat
	}
	tests := []struct {
		name string
		args args
		want *BigFloat
	}{
		{
			name: "1 + i",
			args: args{a: NewBigFloat(1, 0), b: NewBigFloat(0, 1)},
			want: NewBigFloat(1, 1),
		},
		{
			name: "(1 + i) + (1 + i)",
			args: args{a: NewBigFloat(1, 1), b: NewBigFloat(1, 1)},
			want: NewBigFloat(2, 2),
		},
		{
			name: "(3 + 4i) + (5 + 7i)",
			args: args{a: NewBigFloat(3, 4), b: NewBigFloat(5, 7)},
			want: NewBigFloat(8, 11),
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

func TestBigFloatMul(t *testing.T) {
	type args struct {
		a *BigFloat
		b *BigFloat
	}
	tests := []struct {
		name string
		args args
		want *BigFloat
	}{
		{
			name: "1·i",
			args: args{a: NewBigFloat(1, 0), b: NewBigFloat(0, 1)},
			want: NewBigFloat(0, 1),
		},
		{
			name: "(1 + i)(1 + i)",
			args: args{a: NewBigFloat(1, 1), b: NewBigFloat(1, 1)},
			want: NewBigFloat(0, 2),
		},
		{
			name: "(2 + i)(-1 + 3i)",
			args: args{a: NewBigFloat(2, 1), b: NewBigFloat(-1, 3)},
			want: NewBigFloat(-5, 5),
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

func TestBigFloatAbs(t *testing.T) {
	type args struct {
		z *BigFloat
	}
	tests := []struct {
		name string
		args args
		want *big.Float
	}{
		{
			name: "|2|",
			args: args{z: NewBigFloat(2, 0)},
			want: big.NewFloat(2),
		},
		{
			name: "|2i|",
			args: args{z: NewBigFloat(0, 2)},
			want: big.NewFloat(2),
		},
		{
			name: "|3 + 4i|",
			args: args{z: NewBigFloat(3, 4)},
			want: big.NewFloat(5),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.z.Abs(); got.Cmp(tt.want) != 0 {
				t.Errorf("%v.Abs() = %v, want %v", tt.args.z, got, tt.want)
			}
		})
	}
}

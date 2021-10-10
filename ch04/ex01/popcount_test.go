package popcount

import (
	"crypto/sha256"
	"testing"
)

func TestPopCount(t *testing.T) {
	type args struct {
		b byte
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "0x00",
			args: args{b: 0x00},
			want: 0,
		},
		{
			name: "0x63",
			args: args{b: 0x63},
			want: 4,
		},
		{
			name: "0xff",
			args: args{b: 0xff},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PopCount(tt.args.b); got != tt.want {
				t.Errorf("PopCount(%v) = %v, want %v", tt.args.b, got, tt.want)
			}
		})
	}
}

func TestPopCountDiff(t *testing.T) {
	type args struct {
		b1 [32]byte
		b2 [32]byte
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "sha256(\"x\") - sha256(\"x\")",
			args: args{b1: sha256.Sum256([]byte("x")), b2: sha256.Sum256([]byte("x"))},
			want: 0,
		},
		{
			name: "sha256(\"x\") - sha256(\"X\")",
			args: args{b1: sha256.Sum256([]byte("x")), b2: sha256.Sum256([]byte("X"))},
			want: 125,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PopCountDiff(tt.args.b1, tt.args.b2); got != tt.want {
				t.Errorf("PopCountDiff(%v, %v) = %v, want %v", tt.args.b1, tt.args.b2, got, tt.want)
			}
		})
	}
}

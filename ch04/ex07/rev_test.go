package rev

import (
	"bytes"
	"testing"
)

func TestReverseUTF8(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "empty string",
			args: args{[]byte("")},
			want: []byte(""),
		},
		{
			name: "single ascii character",
			args: args{[]byte("x")},
			want: []byte("x"),
		},
		{
			name: "multiple ascii characters",
			args: args{[]byte("abcdefgABCDEFG")},
			want: []byte("GFEDCBAgfedcba"),
		},
		{
			name: "single mathematical script",
			args: args{[]byte("𝒳")},
			want: []byte("𝒳"),
		},
		{
			name: "multiple mathematical scripts",
			args: args{[]byte("𝒳𝒴𝒵𝓍𝓎𝓏")},
			want: []byte("𝓏𝓎𝓍𝒵𝒴𝒳"),
		},
		{
			name: "single Japanese characters",
			args: args{[]byte("あ")},
			want: []byte("あ"),
		},
		{
			name: "multiple Japanese characters",
			args: args{[]byte("あいうえお")},
			want: []byte("おえういあ"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := make([]byte, len(tt.args.b))
			copy(got, tt.args.b)
			if reverseUTF8(got); !bytes.Equal(got, tt.want) {
				t.Errorf("reverseUTF8(%v) = %v, want %v", tt.args.b, got, tt.want)
			}
		})
	}
}

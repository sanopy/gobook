package compress

import (
	"bytes"
	"testing"
)

func TestCompressSpaces(t *testing.T) {
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
			name: "ascii spaces",
			args: args{[]byte("a bc     de   f    g")},
			want: []byte("a bc de f g"),
		},
		{
			name: "ascii whitespaces",
			args: args{[]byte("a\tbc\n\vde\r   \nfg")},
			want: []byte("a bc de fg"),
		},
		{
			name: "multibyte spaces",
			args: args{[]byte("a　　b　cd　　e　　fg")},
			want: []byte("a b cd e fg"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := make([]byte, len(tt.args.b))
			copy(got, tt.args.b)
			if got = compressSpaces(got); !bytes.Equal(got, tt.want) {
				t.Errorf("compressSpaces(%v) = %v, want %v", tt.args.b, got, tt.want)
			}
		})
	}
}

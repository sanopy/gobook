package expand

import (
	"strings"
	"testing"
)

func TestExpand(t *testing.T) {
	type args struct {
		s string
		f func(string) string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "s: \"\", f: nil",
			args: args{"", nil},
			want: "",
		},
		{
			name: "s: \"a $bc de $f g\", f: nil",
			args: args{"a $bc de $f g", nil},
			want: "a $bc de $f g",
		},
		{
			name: "s: \"\", f: simply return s",
			args: args{"", func(s string) string {
				return s
			}},
			want: "",
		},
		{
			name: "s: \"a $bc de $f g\", f: simply return s",
			args: args{"a $bc de $f g", func(s string) string {
				return s
			}},
			want: "a bc de f g",
		},
		{
			name: "s: \"a $bc de $f g\", f: strings.ToUpper",
			args: args{"a $bc de $f g", strings.ToUpper},
			want: "a BC de F g",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := expand(tt.args.s, tt.args.f); got != tt.want {
				t.Errorf("expand(%v, f) = %v, want %v", tt.args.s, got, tt.want)
			}
		})
	}
}

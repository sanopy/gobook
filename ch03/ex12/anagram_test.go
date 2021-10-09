package anagram

import "testing"

func TestCommaPositiveNumber(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "empty string",
			args: args{s1: "", s2: ""},
			want: true,
		},
		{
			name: "ascii: tar - rat",
			args: args{s1: "tar", s2: "rat"},
			want: true,
		},
		{
			name: "ascii: cat - rat",
			args: args{s1: "cat", s2: "rat"},
			want: false,
		},
		{
			name: "multi byte: バージョン - バンジョー",
			args: args{s1: "バージョン", s2: "バンジョー"},
			want: true,
		},
		{
			name: "multi byte: バージョン - バンチョウ",
			args: args{s1: "バージョン", s2: "バンチョウ"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isAnagram(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("isAnagram(%v, %v) = %v, want %v", tt.args.s1, tt.args.s2, got, tt.want)
			}
		})
	}
}

package palindrome

import "testing"

type Runes []rune

func (b Runes) Less(i, j int) bool {
	return b[i] < b[j]
}

func (b Runes) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func (b Runes) Len() int {
	return len(b)
}

func TestIsPalindrome(t *testing.T) {
	type args struct {
		r []rune
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "empty string",
			args: args{[]rune("")},
			want: true,
		},
		{
			name: "single ascii character",
			args: args{[]rune("a")},
			want: true,
		},
		{
			name: "single Japanese character",
			args: args{[]rune("あ")},
			want: true,
		},
		{
			name: "even number of ascii characters: aa",
			args: args{[]rune("aa")},
			want: true,
		},
		{
			name: "even number of ascii characters: ab",
			args: args{[]rune("ab")},
			want: false,
		},
		{
			name: "even number of ascii characters: renner",
			args: args{[]rune("renner")},
			want: true,
		},
		{
			name: "even number of ascii characters: render",
			args: args{[]rune("render")},
			want: false,
		},
		{
			name: "even number of Japanese characters: ああ",
			args: args{[]rune("ああ")},
			want: true,
		},
		{
			name: "even number of Japanese characters: あい",
			args: args{[]rune("あい")},
			want: false,
		},
		{
			name: "even number of Japanese characters: かるいいるか",
			args: args{[]rune("かるいいるか")},
			want: true,
		},
		{
			name: "even number of Japanese characters: かるいすいか",
			args: args{[]rune("かるいすいか")},
			want: false,
		},
		{
			name: "odd number of ascii characters: aba",
			args: args{[]rune("aba")},
			want: true,
		},
		{
			name: "odd number of ascii characters: abc",
			args: args{[]rune("abc")},
			want: false,
		},
		{
			name: "odd number of ascii characters: racecar",
			args: args{[]rune("racecar")},
			want: true,
		},
		{
			name: "odd number of ascii characters: pacecar",
			args: args{[]rune("pacecar")},
			want: false,
		},
		{
			name: "odd number of Japanese characters: やおや",
			args: args{[]rune("やおや")},
			want: true,
		},
		{
			name: "odd number of Japanese characters: やしき",
			args: args{[]rune("やしき")},
			want: false,
		},
		{
			name: "odd number of Japanese characters: たけやぶやけた",
			args: args{[]rune("たけやぶやけた")},
			want: true,
		},
		{
			name: "odd number of Japanese characters: たけやぶよけた",
			args: args{[]rune("たけやぶよけた")},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPalindrome(Runes(tt.args.r)); got != tt.want {
				t.Errorf("IsPalindrome(%v) = %v, want %v", tt.args.r, got, tt.want)
			}
		})
	}
}

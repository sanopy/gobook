package mystrings

import "testing"

func TestJoin(t *testing.T) {
	type args struct {
		sep  string
		elms []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: `sep: ",", elms: {}`,
			args: args{",", []string{}},
			want: "",
		},
		{
			name: `sep: ",", elms: {"Hello"}`,
			args: args{",", []string{"Hello"}},
			want: "Hello",
		},
		{
			name: `sep: ",", elms: {"Hello", "World!"}`,
			args: args{",", []string{"Hello", "World!"}},
			want: "Hello,World!",
		},
		{
			name: `sep: ", ", elms: {"A", "B", "C", "D"}`,
			args: args{", ", []string{"A", "B", "C", "D"}},
			want: "A, B, C, D",
		},
		{
			name: `sep: "", elms: {"A", "B", "C", "D"}`,
			args: args{"", []string{"A", "B", "C", "D"}},
			want: "ABCD",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Join(tt.args.sep, tt.args.elms...); got != tt.want {
				t.Errorf("Join(%v, %v) = %v, want %v", tt.args.sep, tt.args.elms, got, tt.want)
			}
		})
	}
}

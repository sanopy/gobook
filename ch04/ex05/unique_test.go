package unique

import "testing"

func Equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func TestUnique(t *testing.T) {
	type args struct {
		s []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "[]",
			args: args{[]string{}},
			want: []string{},
		},
		{
			name: "[a]",
			args: args{[]string{"a"}},
			want: []string{"a"},
		},
		{
			name: "[a a]",
			args: args{[]string{"a", "a"}},
			want: []string{"a"},
		},
		{
			name: "[a a a]",
			args: args{[]string{"a", "a", "a"}},
			want: []string{"a"},
		},
		{
			name: "[a a b b b a]",
			args: args{[]string{"a", "a", "b", "b", "b", "a"}},
			want: []string{"a", "b", "a"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := make([]string, len(tt.args.s))
			copy(got, tt.args.s)
			if got = unique(got); !Equal(got, tt.want) {
				t.Errorf("unique(%v) = %v, want %v", tt.args.s, got, tt.want)
			}
		})
	}
}

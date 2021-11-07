package counter

import "testing"

func TestWordCounter(t *testing.T) {
	type args struct {
		p []byte
	}
	tests := []struct {
		name string
		args args
		want WordCounter
	}{
		{
			name: "empty string",
			args: args{[]byte("")},
			want: WordCounter(0),
		},
		{
			name: "single word",
			args: args{[]byte("word")},
			want: WordCounter(1),
		},
		{
			name: "multiple words",
			args: args{[]byte("This is a test sentence.")},
			want: WordCounter(5),
		},
		{
			name: "multiple lines",
			args: args{[]byte("This\nis\na\ntest\nsentence.")},
			want: WordCounter(5),
		},
		{
			name: "multiple words and multiple lines",
			args: args{[]byte("This is\na test\nsentence.")},
			want: WordCounter(5),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var c WordCounter
			if c.Write(tt.args.p); c != tt.want {
				t.Errorf("actual: %v, want: %v", c, tt.want)
			}
		})
	}
}

func TestLineCounter(t *testing.T) {
	type args struct {
		p []byte
	}
	tests := []struct {
		name string
		args args
		want LineCounter
	}{
		{
			name: "empty string",
			args: args{[]byte("")},
			want: LineCounter(0),
		},
		{
			name: "single word",
			args: args{[]byte("word")},
			want: LineCounter(1),
		},
		{
			name: "multiple words",
			args: args{[]byte("This is a test sentence.")},
			want: LineCounter(1),
		},
		{
			name: "multiple lines",
			args: args{[]byte("This\nis\na\ntest\nsentence.")},
			want: LineCounter(5),
		},
		{
			name: "multiple words and multiple lines",
			args: args{[]byte("This is\na test\nsentence.")},
			want: LineCounter(3),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var c LineCounter
			if c.Write(tt.args.p); c != tt.want {
				t.Errorf("actual: %v, want: %v", c, tt.want)
			}
		})
	}
}

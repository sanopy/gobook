package comma

import "testing"

func TestComma(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "empty string",
			args: args{s: ""},
			want: "",
		},
		{
			name: "1",
			args: args{s: "1"},
			want: "1",
		},
		{
			name: "12",
			args: args{s: "12"},
			want: "12",
		},
		{
			name: "123",
			args: args{s: "123"},
			want: "123",
		},
		{
			name: "1234",
			args: args{s: "1234"},
			want: "1,234",
		},
		{
			name: "12345",
			args: args{s: "12345"},
			want: "12,345",
		},
		{
			name: "123456",
			args: args{s: "123456"},
			want: "123,456",
		},
		{
			name: "1234567",
			args: args{s: "1234567"},
			want: "1,234,567",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := comma(tt.args.s); got != tt.want {
				t.Errorf("comma(%v) = %v, want %v", tt.args.s, got, tt.want)
			}
		})
	}
}

func TestCommaBuf(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "empty string",
			args: args{s: ""},
			want: "",
		},
		{
			name: "1",
			args: args{s: "1"},
			want: "1",
		},
		{
			name: "12",
			args: args{s: "12"},
			want: "12",
		},
		{
			name: "123",
			args: args{s: "123"},
			want: "123",
		},
		{
			name: "1234",
			args: args{s: "1234"},
			want: "1,234",
		},
		{
			name: "12345",
			args: args{s: "12345"},
			want: "12,345",
		},
		{
			name: "123456",
			args: args{s: "123456"},
			want: "123,456",
		},
		{
			name: "1234567",
			args: args{s: "1234567"},
			want: "1,234,567",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := commaBuf(tt.args.s); got != tt.want {
				t.Errorf("commaBuf(%v) = %v, want %v", tt.args.s, got, tt.want)
			}
		})
	}
}

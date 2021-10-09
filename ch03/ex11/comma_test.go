package comma

import "testing"

func TestCommaPositiveNumber(t *testing.T) {
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
			name: "123456789",
			args: args{s: "123456789"},
			want: "123,456,789",
		},
		{
			name: "1.4589",
			args: args{s: "1.4589"},
			want: "1.4589",
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

func TestCommaNegativeNumber(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "-1",
			args: args{s: "-1"},
			want: "-1",
		},
		{
			name: "-12",
			args: args{s: "-12"},
			want: "-12",
		},
		{
			name: "-123",
			args: args{s: "-123"},
			want: "-123",
		},
		{
			name: "-1234",
			args: args{s: "-1234"},
			want: "-1,234",
		},
		{
			name: "-123456789",
			args: args{s: "-123456789"},
			want: "-123,456,789",
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

func TestCommaDecimalNumber(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "12.4589",
			args: args{s: "12.4589"},
			want: "12.4589",
		},
		{
			name: "123.4589",
			args: args{s: "123.4589"},
			want: "123.4589",
		},
		{
			name: "1234.4589",
			args: args{s: "1234.4589"},
			want: "1,234.4589",
		},
		{
			name: "123456789.4589",
			args: args{s: "123456789.4589"},
			want: "123,456,789.4589",
		},
		{
			name: "-1.4589",
			args: args{s: "-1.4589"},
			want: "-1.4589",
		},
		{
			name: "-12.4589",
			args: args{s: "-12.4589"},
			want: "-12.4589",
		},
		{
			name: "-123.4589",
			args: args{s: "-123.4589"},
			want: "-123.4589",
		},
		{
			name: "-1234.4589",
			args: args{s: "-1234.4589"},
			want: "-1,234.4589",
		},
		{
			name: "-123456789.4589",
			args: args{s: "-123456789.4589"},
			want: "-123,456,789.4589",
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

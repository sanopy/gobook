package params

import (
	"testing"
)

type Data struct {
	Labels     []string `http:"l"`
	MaxResults int      `http:"max"`
	Exact      bool     `http:"x"`
	Str        string   `http:"s"`
	Float      float64  `http:"f"`
}

func TestPack(t *testing.T) {
	tests := []struct {
		data Data
		want string
	}{
		{
			data: Data{},
			want: "?max=0&x=false&f=0",
		},
		{
			data: Data{Labels: []string{"Hello", "World"}},
			want: "?l=Hello&l=World&max=0&x=false&f=0",
		},
		{
			data: Data{MaxResults: 100},
			want: "?max=100&x=false&f=0",
		},
		{
			data: Data{Exact: true},
			want: "?max=0&x=true&f=0",
		},
		{
			data: Data{Str: "Hello"},
			want: "?max=0&x=false&s=Hello&f=0",
		},
		{
			data: Data{Float: 0.15},
			want: "?max=0&x=false&f=0.15",
		},
	}
	for _, tt := range tests {
		if got := Pack(&tt.data); got != tt.want {
			t.Errorf("Pack(%v) returned %v, want %v", tt.data, got, tt.want)
		}
	}
}

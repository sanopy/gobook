package sexpr

import (
	"bytes"
	"reflect"
	"strings"
	"testing"
)

type Movie struct {
	Title, Subtitle string
	Year            int
	Color           bool
	Actor           map[string]string
	Oscars          []string
	Sequel          *string
	E               float32
	Pi              float64
	Comp64          complex64
	Comp128         complex128
	Interface       interface{}
}

var strangelove = Movie{
	Title:    "Dr. Strangelove",
	Subtitle: "How I Learned to Stop Worrying and Love the Bomb",
	Year:     1964,
	Color:    false,
	Actor: map[string]string{
		"Dr. Strangelove":            "Peter Sellers",
		"Grp. Capt. Lionel Mandrake": "Peter Sellers",
		"Pres. Merkin Muffley":       "Peter Sellers",
		"Gen. Buck Turgidson":        "George C. Scott",
		"Brig. Gen. Jack D. Ripper":  "Sterling Hayden",
		`Maj. T.J. "King" Kong`:      "Slim Pickens",
	},
	Oscars: []string{
		"Best Actor (Nomin.)",
		"Best Adapted Screenplay (Nomin.)",
		"Best Director (Nomin.)",
		"Best Picture (Nomin.)",
	},
	E:         2.718281828459045235360287471352,
	Pi:        3.14159265358979323846264338327950288,
	Comp64:    complex(1.01, 2.05),
	Comp128:   complex(12.22, 24.33),
	Interface: 5,
	// Interface: []int{1, 2, 3},
}

// Test verifies that encoding and decoding a complex data value
// produces an equal result.
//
// The test does not make direct assertions about the encoded output
// because the output depends on map iteration order, which is
// nondeterministic.  The output of the t.Log statements can be
// inspected by running the test with the -v flag:
//
// 	$ go test -v gopl.io/ch12/sexpr
//
func Test(t *testing.T) {
	// Encode it
	data, err := Marshal(strangelove)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}
	t.Logf("Marshal() = %s\n", data)

	// Decode it
	var movie Movie
	if err := Unmarshal(data, &movie); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	t.Logf("Unmarshal() = %+v\n", movie)

	// Check equality.
	if !reflect.DeepEqual(movie, strangelove) {
		t.Fatal("not equal")
	}

	// Pretty-print it:
	data, err = MarshalIndent(strangelove)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("MarshalIdent() = %s\n", data)
}

func TestZeroValue(t *testing.T) {
	strangelove := Movie{}

	// Encode it
	data, err := Marshal(strangelove)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}
	t.Logf("Marshal() = %s\n", data)

	// Decode it
	var movie Movie
	if err := Unmarshal(data, &movie); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	t.Logf("Unmarshal() = %+v\n", movie)

	// Check equality.
	if !reflect.DeepEqual(movie, strangelove) {
		t.Fatal("not equal")
	}
}

func TestStream(t *testing.T) {
	// Encode it
	var buf bytes.Buffer
	enc := NewEncoder(&buf)
	if err := enc.Encode(strangelove); err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}
	t.Logf("Marshal() = %s\n", buf.String())

	// Decode it
	var movie Movie
	dec := NewDecoder(&buf)
	if err := dec.Decode(&movie); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	t.Logf("Unmarshal() = %+v\n", movie)

	// Check equality.
	if !reflect.DeepEqual(movie, strangelove) {
		t.Fatal("not equal")
	}
}

func TestDecoderToken(t *testing.T) {
	tests := []struct {
		sexpr string
		want  []Token
	}{
		{
			sexpr: "",
			want:  []Token{},
		},
		{
			sexpr: "()",
			want:  []Token{StartList{}, EndList{}},
		},
		{
			sexpr: "(a b c)",
			want:  []Token{StartList{}, Symbol{"a"}, Symbol{"b"}, Symbol{"c"}, EndList{}},
		},
		{
			sexpr: `("Hello" "World")`,
			want:  []Token{StartList{}, String{"Hello"}, String{"World"}, EndList{}},
		},
		{
			sexpr: "(1 2 3)",
			want:  []Token{StartList{}, Int{1}, Int{2}, Int{3}, EndList{}},
		},
	}
	for _, tt := range tests {
		r := strings.NewReader(tt.sexpr)
		d := NewDecoder(r)
		for _, token := range tt.want {
			if got := d.Token(); !reflect.DeepEqual(got, token) {
				t.Errorf("d.Token() returned %T, want %T", got, token)
			}
		}
	}
}

package sexpr

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestJsonencode(t *testing.T) {
	// Encode it
	data, err := JsonMarshal(strangelove)
	if err != nil {
		t.Fatalf("JsonMarshal failed: %v", err)
	}
	t.Logf("JsonMarshal() = %s\n", data)

	// Decode it
	var movie Movie
	if err := json.Unmarshal(data, &movie); err != nil {
		t.Fatalf("json.Unmarshal failed: %v", err)
	}
	t.Logf("json.Unmarshal() = %+v\n", movie)

	// Check equality.
	if !reflect.DeepEqual(movie, strangelove) {
		t.Fatal("not equal")
	}
}

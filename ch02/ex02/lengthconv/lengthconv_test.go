package lengthconv

import (
	"math"
	"testing"
)

const EPSILON = 0.000001

func TestFToM(t *testing.T) {
	actual := FToM(OneMeterInFeet)
	expected := Meter(1)
	if math.Abs(float64(actual-expected)) >= EPSILON { // actual != expected
		t.Errorf(`FToM(%v) = %v`, OneMeterInFeet, actual)
	}
}

func TestMToF(t *testing.T) {
	actual := MToF(OneFeetInMeters)
	expected := Feet(1)
	if math.Abs(float64(actual-expected)) >= EPSILON { // actual != expected
		t.Errorf(`MToF(%v) = %v`, OneFeetInMeters, actual)
	}
}

package tempconv

import (
	"math"
	"testing"
)

const EPSILON = 0.000001

func TestCToF(t *testing.T) {
	actual := CToF(FreezingC)
	expected := FreezingF
	if math.Abs(float64(actual-expected)) >= EPSILON { // actual != expected
		t.Errorf(`CToF(%v) = %v`, FreezingC, actual)
	}
}

func TestFToC(t *testing.T) {
	actual := FToC(FreezingF)
	expected := FreezingC
	if math.Abs(float64(actual-expected)) >= EPSILON { // actual != expected
		t.Errorf(`FToC(%v) = %v`, FreezingF, actual)
	}
}

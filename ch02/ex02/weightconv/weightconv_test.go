package weightconv

import (
	"math"
	"testing"
)

const EPSILON = 0.000001

func TestPToKg(t *testing.T) {
	actual := PToKg(OneKilogramInPound)
	expected := Kilogram(1)
	if math.Abs(float64(actual-expected)) >= EPSILON { // actual != expected
		t.Errorf(`PToKg(%v) = %v`, OneKilogramInPound, actual)
	}
}

func TestKgToP(t *testing.T) {
	actual := KgToP(OnePoundInKilogram)
	expected := Pound(1)
	if math.Abs(float64(actual-expected)) >= EPSILON { // actual != expected
		t.Errorf(`KgToP(%v) = %v`, OnePoundInKilogram, actual)
	}
}

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

func TestCToK(t *testing.T) {
	actual := CToK(FreezingC)
	expected := FreezingK
	if math.Abs(float64(actual-expected)) >= EPSILON { // actual != expected
		t.Errorf(`CToK(%v) = %v`, FreezingC, actual)
	}
}

func TestFToC(t *testing.T) {
	actual := FToC(FreezingF)
	expected := FreezingC
	if math.Abs(float64(actual-expected)) >= EPSILON { // actual != expected
		t.Errorf(`FToC(%v) = %v`, FreezingF, actual)
	}
}

func TestFToK(t *testing.T) {
	actual := FToK(FreezingF)
	expected := FreezingK
	if math.Abs(float64(actual-expected)) >= EPSILON { // actual != expected
		t.Errorf(`FToK(%v) = %v`, FreezingF, actual)
	}
}

func TestKToC(t *testing.T) {
	actual := KToC(FreezingK)
	expected := FreezingC
	if math.Abs(float64(actual-expected)) >= EPSILON { // actual != expected
		t.Errorf(`KToC(%v) = %v`, FreezingK, actual)
	}
}

func TestKToF(t *testing.T) {
	actual := KToF(FreezingK)
	expected := FreezingF
	if math.Abs(float64(actual-expected)) >= EPSILON { // actual != expected
		t.Errorf(`KToF(%v) = %v`, FreezingK, actual)
	}
}

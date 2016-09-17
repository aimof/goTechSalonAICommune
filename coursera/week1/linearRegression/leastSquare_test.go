package main

import (
	"testing"
)

func TestDiferentSquare(t *testing.T) {
	actual := differentSquare(0, 0, 0, 1)
	expected := 0.0
	if actual != expected {
		t.Errorf("一致しません\n%f", actual)
	}
}

func TestSumDifferentSquare(t *testing.T) {
	point0 := point{0, 0}
	point1 := point{1, 1}
	var testPoints []point
	testPoints = append(testPoints, point0)
	testPoints = append(testPoints, point1)
	actual := sumDifferentSquare(testPoints, 0.0, 1.0)
	expected := 0.0
	if actual != expected {
		t.Errorf("一致しません\n%f", actual)
	}
}

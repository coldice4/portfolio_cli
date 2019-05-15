package main

import "testing"

func TestCalculateFactor(t *testing.T) {
	testString := "231.443"
	factor := calculateFactor(testString)
	if factor != 1000 {
		t.Fail()
	}
}

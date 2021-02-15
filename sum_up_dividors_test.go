package main

import "testing"

func TestSumUpDividors(t *testing.T) {
	res := SumUpDividors(30)
	if res != 72 {
		t.Errorf("Expected 72, but not")
	}
}

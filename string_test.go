package main

import "testing"

func TestCountDuplicate(t *testing.T) {
	s := "foobarbaz"
	count := CountDuplicate(s)
	if count != 3 {
		t.Errorf("Expected 3")
	}
}

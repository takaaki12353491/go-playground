package main

import "testing"

func TestReverse(t *testing.T) {
	s := "dowango"
	res := Reverse(s)
	if res != "ognawod" {
		t.Errorf("Expected ognawod")
	}
}

func TestCountDuplicate(t *testing.T) {
	s := "foobarbaz"
	count := CountDuplicate(s)
	if count != 3 {
		t.Errorf("Expected 3")
	}
}

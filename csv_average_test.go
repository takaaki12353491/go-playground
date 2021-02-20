package main

import (
	"strings"
	"testing"
)

func TestCSVAverage(t *testing.T) {
	s := strings.Join([]string{
		"Math,English",
		"90,70",
		"80,90",
		"70,60",
		"60,30",
	}, "\n")
	res := CSVAverage(s)
	math := res["Math"]
	english := res["English"]
	if math != 75 || english != 62.5 {
		t.Errorf("Math: %f, English: %f", math, english)
	}
}

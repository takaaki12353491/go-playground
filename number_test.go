package main

import (
	"strings"
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	res := FizzBuzz(15)
	FizzBuzz := strings.Count(res, "FizzBuzz")
	Fizz := strings.Count(res, "Fizz") - FizzBuzz
	Buzz := strings.Count(res, "Buzz") - FizzBuzz
	if FizzBuzz != 1 || Fizz != 4 || Buzz != 2 {
		t.Errorf("Expected Fizz:%d, Buzz:%d, FizzBuzz:%d", 4, 2, 1)
	}
}

func TestFibonacci(t *testing.T) {
	res := Fibonacci(10)
	ex := 55
	if res != ex {
		t.Errorf("Expected %d but %d", ex, res)
	}
}

func TestCountAppearance(t *testing.T) {
	res := CountAppearance(100, 7)
	ex := 20
	if res != ex {
		t.Errorf("Expected %d but %d", ex, res)
	}
}

func TestSumUpDividors(t *testing.T) {
	res := SumUpDividors(30)
	ex := 72
	if res != ex {
		t.Errorf("Expected %d but %d", ex, res)
	}
}

func TestMultiplyToMax(t *testing.T) {
	// Missing elements
	nums := []int{1, 1}
	res := MultiplyToMax(nums, 3)
	ex := -1
	if res != ex {
		t.Errorf("Expected %d but %d", ex, res)
	}
	// more than 3
	nums = []int{3, 1, 2, 5, 4}
	res = MultiplyToMax(nums, 3)
	ex = 60
	if res != ex {
		t.Errorf("Expected %d but %d", ex, res)
	}
}

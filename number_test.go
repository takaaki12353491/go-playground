package main

import (
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	f := FizzBuzz(3, 5)
	res := f(1)
	ex := "1"
	if res != ex {
		t.Errorf("Expected %s, but %s", ex, res)
	}
	res = f(3)
	ex = "Fizz"
	if res != ex {
		t.Errorf("Expected %s, but %s", ex, res)
	}
	res = f(5)
	ex = "Buzz"
	if res != ex {
		t.Errorf("Expected %s, but %s", ex, res)
	}
	res = f(15)
	ex = "FizzBuzz"
	if res != ex {
		t.Errorf("Expected %s, but %s", ex, res)
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

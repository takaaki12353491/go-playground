package main

import (
	"math"
)

func SumUpDividors(num int) int {
	sum := 0
	sqrt := math.Sqrt(float64(num))
	for i := 1; i <= int(sqrt); i++ {
		m := num % i
		if m != 0 {
			continue
		}
		d := num / i
		sum += i + d
	}
	return sum
}

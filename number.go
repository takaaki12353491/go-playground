package main

import (
	"math"
	"sort"
	"strconv"
	"strings"
)

func FizzBuzz(num int) string {
	res := ""
	for i := 1; i <= num; i++ {
		d3 := i % 3
		d5 := i % 5
		if d3 == 0 && d5 == 0 {
			res += "FizzBuzz"
		} else if d3 == 0 {
			res += "Fizz"
		} else if d5 == 0 {
			res += "Buzz"
		} else {
			res += strconv.Itoa(i)
		}
		res += "\n"
	}
	return res
}

func Fibonacci(num int) int {
	res := 0
	i1 := 1
	i2 := 1
	switch num {
	case 1:
		res = i1
	case 2:
		res = i2
	default:
		for i := 0; i < num-2; i++ {
			res = i1 + i2
			i1 = i2
			i2 = res
		}
	}
	return res
}

func CountAppearance(num, target int) int {
	count := 0
	for i := 1; i <= num; i++ {
		num := strconv.Itoa(i)
		slice := strings.Split(num, "")
		t := strconv.Itoa(target)
		for _, v := range slice {
			if v == t {
				count++
			}
		}
	}
	return count
}

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

func MultiplyToMax(nums []int, num int) int {
	if len(nums) < num {
		return -1
	}
	sort.Ints(nums)
	l := len(nums)
	res := 1
	for i := 1; i <= num; i++ {
		res *= nums[l-i]
	}
	return res
}

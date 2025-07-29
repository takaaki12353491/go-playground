package main

import (
	"fmt"
	"iter"
)

func loopWithIter() {
	for i := range iterator() {
		fmt.Println(i)
	}
}

func iterator() iter.Seq[int] {
	return func(yield func(int) bool) {
		for i := 1; i <= 10; i++ {
			if !yield(i) {
				break
			}
		}
	}
}

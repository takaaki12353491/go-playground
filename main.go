package main

import (
	"bufio"
	"fmt"
	"os"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	var s string
	if sc.Scan() {
		s = sc.Text()
	}
	fmt.Println(s)
}

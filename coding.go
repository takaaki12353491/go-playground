package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
問題1
1から100までの数字を順番に出力する関数を作成してください。
ただし、出力する数字が3の倍数の場合には「Fizz」を、
出力する数字が5の倍数の場合には「Buzz」を、
出力 する数字が3の倍数かつ5の倍数の場合には「FizzBuzz」を、
それぞれ数字の代わりに出力してください。
また、各数字もしくは文字列の出力の後に改行コード(LF - \n)を出力してください。
問題2
1から1000までのフィボナッチ数列を出力する関数を作成してください。
フィボナッチ数列の定義は下記の通りである。

フィボナッチ数列とは、最初の二つの数字が1で、三つ目以降がすべて直前の二つの数字の和になっている数列。
例) 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89…
問題3
0から9999までの数字を順番に1刻みで出力するプログラムを作成しました。
このプログラムを実行した際に、数字の7は何回出力されますか?
例えば、7777が出力された場合、数字の7は4回出力されたものとしてカウントします。
この問題の答え、もしくは答えを計算するためのプログラムを解答してください。
*/

func problem1() {
	for i := 1; i <= 100; i++ {
		d3 := i % 3
		d5 := i % 5
		if d3 == 0 && d5 == 0 {
			fmt.Println("FizzBuzz")
		} else if d3 == 0 {
			fmt.Println("Fizz")
		} else if d5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}

func problem2() {
	i1 := 0
	i2 := 1
	for {
		res := i1 + i2
		fmt.Println(res)
		if res > 1000 {
			break
		}
		i1 = i2
		i2 = res
	}
}

func problem3() {
	count := 0
	for i := 0; i < 10000; i++ {
		num := strconv.Itoa(i)
		slice := strings.Split(num, "")
		for _, v := range slice {
			if strings.Contains(v, "7") {
				count++
			}
		}
	}
	fmt.Println(count)
}

package main

import (
	"fmt"
	"sort"
	"strings"
)

func Reverse(s string) string {
	res := ""
	slice := strings.Split(s, "")
	for i := len(slice) - 1; i >= 0; i-- {
		res += slice[i]
	}
	return res
}

func CountDuplicate(s string) int {
	slice := strings.Split(s, "")
	count := 0
	checked := map[string]bool{}
	for _, v := range slice {
		if _, ok := checked[v]; ok {
			count++
			continue
		}
		checked[v] = true
	}
	return count
}

/*
Input: apple orange apple apple oange orange grape apple
Output: {oange 1} {grape 1} {orange 2} {apple 4}
*/

func SortString(s string) {
	slice := strings.Split(s, " ")
	m := map[string]int{}
	for _, v := range slice {
		if _, ok := m[v]; !ok {
			m[v] = 1
			continue
		}
		m[v] += 1
	}
	entries := []Entry{}
	for k, v := range m {
		entry := Entry{k, v}
		entries = append(entries, entry)
	}
	sort.Slice(entries, func(i, j int) bool {
		return entries[i].value < entries[j].value
	})
	fmt.Println(entries)
}

type Entry struct {
	name  string
	value int
}

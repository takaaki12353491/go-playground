package main

import (
	"fmt"
	"sort"
	"strings"
)

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
		m[v] = m[v] + 1
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

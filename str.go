package main

import (
	"fmt"
	"strings"
)

// "apple orange apple apple oange orange grape apple"

func sort() {
	str := "apple orange apple apple oange orange grape apple"
	slice := strings.Split(str, " ")
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

package main

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func ExecOperation(operations []string) []int {
	repo := Repository{
		values: map[string]Data{},
	}
	res := []int{}
	for _, operation := range operations {
		slice := strings.Split(operation, " ")
		switch slice[0] {
		case "add":
			k := slice[1]
			v, err := strconv.Atoi(slice[2])
			if err != nil {
				continue
			}
			repo.Add(k, v)
		case "get":
			k := slice[1]
			v, err := repo.Get(k)
			if err != nil {
				fmt.Println(-1)
				continue
			}
			res = append(res, v)
		case "remove":
			k := slice[1]
			v, err := repo.Remove(k)
			if err != nil {
				fmt.Println(-1)
				continue
			}
			res = append(res, v)
		case "evict":
			repo.Evict()
		case "exit":
			break
		default:
		}
	}
	return res
}

type (
	Repository struct {
		values map[string]Data
	}
	Data struct {
		value int
		turn  int
	}
)

func (repo *Repository) Add(k string, v int) {
	repo.values[k] = Data{
		value: v,
		turn:  0,
	}
	repo.elaspe()
}

func (repo *Repository) Get(k string) (int, error) {
	data, ok := repo.values[k]
	if !ok {
		return data.value, errors.New("Element is not found")
	}
	data.turn = 0
	repo.elaspe()
	return data.value, nil
}

func (repo *Repository) Remove(k string) (int, error) {
	data, ok := repo.values[k]
	if !ok {
		return data.value, errors.New("Element is not found")
	}
	delete(repo.values, k)
	return data.value, nil
}

func (repo *Repository) Evict() {
	entries := []Entry{}
	for k, data := range repo.values {
		entry := Entry{k, data.value}
		entries = append(entries, entry)
	}
	sort.Slice(entries, func(i, j int) bool {
		return entries[i].value < entries[j].value
	})
	lastKey := entries[len(entries)-1].name
	delete(repo.values, lastKey)
}

func (repo *Repository) elaspe() {
	for _, data := range repo.values {
		data.turn += 1
	}
}

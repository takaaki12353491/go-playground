package main

import (
	"encoding/csv"
	"io"
	"log"
	"math"
	"strconv"
	"strings"
)

func CSVAverage(s string) map[string]float64 {
	r := csv.NewReader(strings.NewReader(s))
	titles, err := r.Read()
	if err != nil {
		log.Fatal(err)
	}
	sums := make([]int, len(titles), len(titles))
	length := 0
	for {
		line, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		for k, v := range line {
			score, err := strconv.Atoi(v)
			if err != nil {
				log.Fatal(err)
			}
			sums[k] += score
		}
		length++
	}
	if length == 0 {
		log.Fatal("no record")
	}
	averages := map[string]float64{}
	for k, title := range titles {
		f := float64(sums[k]) / float64(length)
		average := math.Round(f*100) / 100
		if err != nil {
			log.Fatal(err)
		}
		averages[title] = average
	}
	return averages
}

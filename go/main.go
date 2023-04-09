package main

import (
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func sortValues(values []int) []int {
	//sort.Ints(values[:])

	// trying bubble sort for now
	n := len(values)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if values[j] > values[j+1] {
				values[j], values[j+1] = values[j+1], values[j]
			}
		}
	}

	return values
}

func main() {
	values, err := os.ReadFile("../values")
	check(err)

	var formatted []int

	for _, v := range values {
		formatted = append(formatted, int(v))
	}

	var sorted []int = sortValues(formatted)

	fmt.Println(values)
	fmt.Println(sorted)
}

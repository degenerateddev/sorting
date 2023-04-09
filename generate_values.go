package main

import (
	"bufio"
	"math/rand"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	var max int = 100
	var limit int = 10

	f, err := os.Create("values")
	check(err)

	defer f.Close()

	f.Sync()

	w := bufio.NewWriter(f)

	for i := 1; i <= limit; i++ {
		r := rand.Intn(max)
		r, err := w.WriteString(strconv.Itoa(r) + "\n")
		check(err)
	}

	w.Flush()
}

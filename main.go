package main

import (
	"fmt"

	"github.com/jovi8850/go-trimmed-mean/trimmedmean"
)

func main() {
	data := []float64{1, 2, 3, 100, 101}

	result, err := trimmedmean.SymmetricTrimmedMean(data, 0.05)
	if err != nil {
		panic(err)
	}

	fmt.Println("5% symmetric trimmed mean:", result)
}

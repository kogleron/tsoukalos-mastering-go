package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Need at least one float")
		os.Exit(1)
	}

	arguments := os.Args
	min, _ := strconv.ParseFloat(arguments[1], 64)
	max, _ := strconv.ParseFloat(arguments[1], 64)
	maxI := len(arguments)

	for i := 2; i < maxI; i++ {
		n, _ := strconv.ParseFloat(arguments[i], 64)

		if n > max {
			max = n
		} else if n < min {
			min = n
		}
	}

	fmt.Println("Min: ", min)
	fmt.Println("Max: ", max)
}

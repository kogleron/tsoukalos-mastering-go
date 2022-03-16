package main

import "fmt"

func main() {
	a6 := []int{-10, 1, 2, 3, 4, 5}
	a4 := []int{-1, -2, -3, -4}

	fmt.Println()
	fmt.Println("Before")
	fmt.Println("a6:", a6)
	fmt.Println("a4:", a4)

	// The number of copied elements will be the minimum of len(a6) or len(a4)
	copy(a6, a4)

	fmt.Println()
	fmt.Println("After")
	fmt.Println("a6:", a6)
	fmt.Println("a4:", a4)
	fmt.Println()
}

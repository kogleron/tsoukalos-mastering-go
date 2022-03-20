package main

import "fmt"

// iota example
func main() {
	const (
		// iota can be used in expressions
		p2_0 int = 1 << iota
		_        // low dash for skipping iota values
		p2_2
		p2_3
		_
		p2_5
	)

	fmt.Println("2^0=", p2_0)
	fmt.Println("2^2=", p2_2)
	fmt.Println("2^3=", p2_3)
	fmt.Println("2^5=", p2_5)
}

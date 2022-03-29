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

	const (
		p4_0 int = 1 << (iota * 2)
		_
		_
		p4_3
		_
		p4_5
	)

	fmt.Println("2^0=", p2_0)
	fmt.Println("2^2=", p2_2)
	fmt.Println("2^3=", p2_3)
	fmt.Println("2^5=", p2_5)

	fmt.Println("")

	fmt.Println("4^0=", p4_0)
	fmt.Println("4^3=", p4_3)
	fmt.Println("4^5=", p4_5)
}

package main

import "fmt"

// Work with nil maps
func main() {
	var mymap map[string]bool

	fmt.Println("Is nil = ", mymap == nil)

	// It's OK
	val, has := mymap["some"]
	fmt.Println("val:", val, "; has:", has)

	// It's OK
	for n, val := range mymap {
		fmt.Println(n, val)
	}

	// Panic
	mymap["another"] = false
}

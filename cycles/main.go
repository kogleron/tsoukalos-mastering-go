package main

import "fmt"

func usualFor() {
	fmt.Println()
	fmt.Print("Usual: ")
	for i := 0; i < 3; i++ {
		fmt.Print(i, " ")
	}
}

func whileLike() {
	fmt.Println()
	fmt.Print("While like: ")
	i := 0
	for {
		if i == 3 {
			break
		}
		fmt.Print(i, " ")
		i++
	}
}

func doWhileLike() {
	fmt.Println()
	fmt.Print("Do-while like: ")
	flag := true
	i := 0
	for flag {
		fmt.Print(i, " ")
		i++
		flag = i != 3
	}
}

func rangeFor() {
	is := []int{0, 1, 2}
	fmt.Println()
	fmt.Print("Range: ")
	for _, i := range is {
		fmt.Print(i, " ")
	}
}

func main() {
	usualFor()
	whileLike()
	doWhileLike()
	rangeFor()
}

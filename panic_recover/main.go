package main

// Handling panic with recover
import "fmt"

func a() {
	fmt.Println("Inside a()")

	defer func() {
		if c := recover(); c != nil {
			fmt.Println("Recover inside a()!")
		}
	}()

	fmt.Println("About to call b()")
	b()
	fmt.Println("b() existed!")
	fmt.Println("Exiting a()")
}

func b() {
	fmt.Println("Inside b()")
	panic("Panic in b()!")
	fmt.Println("Exiting b()")
}

func main() {
	a()
	fmt.Println("main() ended!")
}

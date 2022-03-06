package main

// #include <stdio.h>
// void callC(){
// 	printf("Calling C code!\n");
// }
import "C"
import "fmt"

//Calls C code from a GO file
func main() {
	fmt.Println("A GO statement!")
	C.callC()
	fmt.Println("Another GO statement!")
}

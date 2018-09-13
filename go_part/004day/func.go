package main

import "fmt"

func main() {
	function1()
}

func function1() {
	fmt.Println("IN function1 at the top\n")
	defer function2()
	fmt.Println("In function1 at the bottom!\n")
}

func function2() {
	fmt.Println("Function2:Defered until the end ofg the calling function!")
}

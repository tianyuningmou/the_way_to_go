package main

import "fmt"

var a string

func main() {
	a = "G"
	fmt.Printf("%p", &a)
	print(a + "\n")
	f1()
}

func f1() {
	a := "O"
	fmt.Printf("%p", &a)
	print(a + "\n")
	f2()
}

func f2() {
	fmt.Printf("%p", &a)
	print(a + "\n")
}

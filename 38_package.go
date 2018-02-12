package main

import (
	"./pack"
	"fmt"
)

func main() {
	var test1 string
	test1 = pack.ReturnStr()
	fmt.Printf("ReturnStr from package1: %s\n", test1)
	fmt.Printf("Integer from package1: %d\n", pack.Pack1Int)
}

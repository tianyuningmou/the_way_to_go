package main

import (
	    "fmt"
	"runtime"
	     "os"
)

func main()  {
	var goos string = runtime.GOOS
	fmt.Printf("The operating system is：%s\n", goos)
	path := os.Getenv("PATH")
	fmt.Printf("Path is %s\n", path)
}

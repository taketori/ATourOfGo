package main

import "fmt"

func main() {
	var a string
	a = "guys"
	defer fmt.Println(a)
	a = "world"

	fmt.Println("hello")
}

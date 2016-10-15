package main

import "fmt"

func main() {
	//v := 42	// int
	//v := 3.142	// float64
	//v := 0.867 + 0.5i	// complex128
	//v := "abcd" // string
	//v := []rune("abcd") // rune []int32
	v := '❤' // rune int32

	fmt.Printf("%v is of type %T\n", v, v)
}

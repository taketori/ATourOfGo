package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	s = s[:0]
	printSlice(s) //=> [], cap=6

	s = s[:4]
	printSlice(s) //=> [2, 3, 5, 7], still cap=6

	s = s[2:]
	printSlice(s) //=> [5, 7],cap=4

	s = s[0:]
	printSlice(s) //=> [5, 7], cap=4

	s = s[1:]
	printSlice(s) //=> [7], cap=3

	s = s[:]
	printSlice(s) //=> [7], cap=3

	s = append(s, 0, 0, 0, 0) //=> 7, 0, 0, 0], cap=6
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)

}

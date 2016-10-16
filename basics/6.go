package main

import "fmt"

func swap(x, y string) (newX string, newY string) {
	newX = y
	newY = x
	return
}

func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}

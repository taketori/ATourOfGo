package main

import "fmt"

func split(sum float32) (x, y float32) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(split(17))
}

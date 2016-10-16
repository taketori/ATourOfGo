package main

import "fmt"

func main() {
	i := 5
	sum := 0
	for i = 0; i < 10; i++ {
		sum += 1
	}
	fmt.Println(sum, i)
}

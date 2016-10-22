package main

import "fmt"

func Sqrt(x float64) float64 {
	r := float64(x)
	for i := 1; i <= 10; i++ {
		r = r - (r*r-x)/(2*r)

	}
	return r
}

func main() {
	fmt.Println(Sqrt(2))
}

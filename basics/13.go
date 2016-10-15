package main

import (
	"fmt"
	"math"
)

func main() {
	x, y := 3, 4
	f := math.Sqrt(x*x + y*y)
	z := f
	fmt.Println(x, y, z)
}

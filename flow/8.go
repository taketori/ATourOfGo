package main

import "fmt"

func Sqrt(x float64) float64 {
	r := float64(x)
	d := float64(0.0001)
	for t := float64(x + d); t-r >= d; {
		t = r
		r = t - (t*t-x)/(2*t)
	}
	return r
}

func main() {
	fmt.Println(Sqrt(3))
}

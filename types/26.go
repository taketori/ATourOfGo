package main

import "fmt"

func fibonacci() func() int {
	f0 := 0
	f1 := 1
	return func() int {
		r := f1 + f0
		f0 = f1
		f1 = r
		return r
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

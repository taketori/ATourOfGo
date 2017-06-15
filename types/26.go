package main

import "fmt"

func fibonacci() func() int {
	f := []int{0, 1}
	i := 0
	return func() int {
		f = append(f, f[len(f)-2]+f[len(f)-1])
		i += 1
		return f[i-1]
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

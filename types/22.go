package main

import "fmt"

func main() {
	k := "Answer"
	m := map[string]int{k: 42}

	//m[k] = 42
	fmt.Println("The value:", m[k])

	m[k] = 48
	fmt.Println("The value:", m[k])

	delete(m, k)
	fmt.Println("The value:", m[k])

	v, ok := m[k]
	fmt.Println("The value:", v, " Present?:", ok)

}

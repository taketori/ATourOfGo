package main

import "fmt"

func do(i interface{}) {
	switch v := i.(type) {
	case nil:
		fmt.Println("i is nil.")
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know aboout type %T", v)
	}
}

func main() {
	do(nil)
	do(21)
	do("hello")
	do(true)
}

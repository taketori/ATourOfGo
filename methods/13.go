package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {

}

func main() {
	var i I
	i = &T{"hello"}
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

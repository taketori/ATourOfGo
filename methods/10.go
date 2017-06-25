package main

import "fmt"

//I ...
type I interface {
	M()
}

//T ...
type T struct {
	S string
}

//M ...
func (t T) M() {
	fmt.Println(t.S)
}

func main() {
	var i I = T{"hello"}
	i.M()

}

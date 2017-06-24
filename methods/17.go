package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years old)", p.Name, p.Age)
}

func main() {
	a := Person{"taketori", -1}
	b := Person{"x10d", 100}
	fmt.Println(a, b)

}

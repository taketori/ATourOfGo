package main

import "fmt"
import "math"

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64) {
	v.X *= f
	v.Y *= f
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func ScaleFunc(v *Vertex, f float64) {
	v.X *= f
	v.Y *= f
}

func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
	fmt.Println(Abs(v))

	v.Scale(2)
	ScaleFunc(&v, 10)

	p := &Vertex{4, 3}
	fmt.Println(p.Abs())
	fmt.Println(Abs(*p))

	p.Scale(3)
	ScaleFunc(p, 8)

	fmt.Println(v, p)
}

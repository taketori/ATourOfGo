package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell labs": Vertex{
		40.68433, -74.39967,
	},
	"Google": Vertex{
		37.42202,
		-122.048408,
	},
}

func main() {
	fmt.Println(m)
}

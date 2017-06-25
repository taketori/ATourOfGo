package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %g", e)
}

func Sqrt(x float64) (float64, error) {
	var r float64 = x
	switch {
	default:
		return 0, ErrNegativeSqrt(x)
	case x < 0:
		return 0, ErrNegativeSqrt(x)
	case x == 0:
		return 0, nil
	case x > 0:
		for i := 0; i < 10; i++ {
			r = r - (math.Pow(r, 2)-x)/(2*r)
		}
		return r, nil
	}

}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}

package main

import (
	"fmt"
	"time"
)

func f() (r time.Weekday) {
	fmt.Println("this will not be called.")
	r = time.Thursday
	return
}

func main() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	default:
		fmt.Println("Too far away.")
	case f():
		fmt.Println("It's Thursday.")
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tommorow")
	case today + 2:
		fmt.Println("In Two days.")
	}
}

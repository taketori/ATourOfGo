package main

import "fmt"

func split(sum int) (x, y int){
  x = sum * 4 / 9
  y = sum - x
  return
  // when sum = 17, this will return both 7 and 10.
}

func main() {
  fmt.Println(split(17))
}
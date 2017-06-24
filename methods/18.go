package main

import "fmt"
import "strconv"
import "strings"

type IPaddr [4]byte

func (ip IPaddr) String() string {
	var r string
	for _, v := range ip {
		r += strconv.Itoa(int(v)) + "."
	}
	return strings.TrimRight(r, ".")
}

func main() {
	hosts := map[string]IPaddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}

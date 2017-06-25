package main

import (
	"io"
	"math"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func rot13(b byte) (r byte) {
	switch {
	case b < 'A' || 'z' < b:
		r = b
	case b <= 'Z':
		r = byte(math.Mod(float64(b-'A'+13), 26)) + 'A'
	case b >= 'a':
		r = byte(math.Mod(float64(b-'a'+13), 26)) + 'a'
	default:
	}
	return
}

func (r13 rot13Reader) Read(p []byte) (n int, err error) {
	n, err = r13.r.Read(p)
	for i := 0; i < n; i++ {
		p[i] = rot13(p[i])
	}
	return
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}

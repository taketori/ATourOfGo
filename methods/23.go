package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func rot13(b byte) (r byte) {
	switch {
	case b < 'A':
		r = b
	case b <= 'M':
		r = b + 13
	case b <= 'Z':
		r = b - 13
	case b > 'z':
		r = b
	case b > 'm':
		r = b - 13
	case b >= 'a':
		r = b + 13
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

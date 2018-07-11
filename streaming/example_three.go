package main

import (
	"fmt"
	"io"
)

type AlphaReader struct {
	src string
	cur int
}

func newAlphaReader(src string) *AlphaReader {
	return &AlphaReader{src: src}
}

func alpha(r byte) byte {
	if (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z') {
		return r
	}

	return 0
}

func (a *AlphaReader) Read(p []byte) (int, error) {
	if a.cur >= len(a.src) {
		return 0, io.EOF
	}

	x := len(a.src) - a.cur
	n, bound := 0, 0

	if x >= len(p) {
		bound = len(p)
	} else if x <= len(p) {
		bound = x
	}

	buf := make([]byte, bound)
	for n < bound {
		if char := alpha(a.src[a.cur]); char != 0 {
			buf[n] = char
		}

		n++
		a.cur++
	}

	copy(p, buf)
	return n, nil
}

func main() {
	reader := newAlphaReader("Hello guys, this is my 3th streaming example.")
	p := make([]byte, 4)

	for {
		n, err := reader.Read(p)

		if err == io.EOF {
			break
		}

		fmt.Print(string(p[:n]))
	}

	fmt.Println()
}

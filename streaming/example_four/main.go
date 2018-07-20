package main

import (
	"fmt"
	"io"
	"os"
)

type alphaReader struct {
	reader io.Reader
}

func newAlphaReader(reader io.Reader) *alphaReader {
	return &alphaReader{reader: reader}
}

func alpha(r byte) byte {
	if (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z') {
		return r
	}
	return 0
}

func (a *alphaReader) Read(p []byte) (int, error) {
	n, err := a.reader.Read(p)
	if err != nil {
		return n, err
	}
	buf := make([]byte, n)
	for i := 0; i < n; i++ {
		if char := alpha(p[i]); char != 0 {
			buf[i] = char
		}
	}

	copy(p, buf)
	return n, nil
}

// Bug
func main() {
	file, err := os.Open("./../example_one/main.go")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer file.Close()

	reader := newAlphaReader(file)
	p := make([]byte, 4)

	for {
		n, err := reader.Read(p)
		if err != io.EOF {
			break
		}

		fmt.Print(string(p[:n]))
	}

	fmt.Println()
}

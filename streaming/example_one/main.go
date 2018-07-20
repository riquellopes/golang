package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	reader := strings.NewReader("First example with streaming")
	p := make([]byte, 4)

	for {
		n, err := reader.Read(p)

		if err == io.EOF {
			break
		}

		fmt.Print(string(p[:n]))
	}

	fmt.Println("")
}

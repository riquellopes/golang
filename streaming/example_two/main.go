package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	reader := strings.NewReader("My second example with streaming.")
	p := make([]byte, 4)

	for {
		n, err := reader.Read(p)
		if err != nil {
			if err != io.EOF {
				fmt.Println(string(p[:n]))
				break
			}

			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println(string(p[:n]))
	}
}

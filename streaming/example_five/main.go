package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	proverbs := []string{
		"Jesus is love",
		"For God so loved the world that",
	}

	var writer bytes.Buffer

	for _, p := range proverbs {
		n, err := writer.Write([]byte(p))

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		if n != len(p) {
			fmt.Println("Failed to write data")
			os.Exit(1)
		}
	}

	fmt.Println(writer.String())
}

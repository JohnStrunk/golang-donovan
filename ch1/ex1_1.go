package main

import (
	"fmt"
	"os"
)

func main() {
	var line, sep string

	for _, word := range os.Args[0:] {
		line += sep + word
		sep = " "
	}

	fmt.Println(line)
}

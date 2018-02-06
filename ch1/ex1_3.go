package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {

	start := time.Now()
	var line, sep string
	for _, word := range os.Args[1:] {
		line += sep + word
		sep = " "
	}
	ns := time.Since(start).Nanoseconds()

	fmt.Println(line, ns)

	start = time.Now()
	line = strings.Join(os.Args[1:], " ")
	ns = time.Since(start).Nanoseconds()
	fmt.Println(line, ns)
}

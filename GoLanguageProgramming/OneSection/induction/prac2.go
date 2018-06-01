package main

import (
	"fmt"
	"os"
)

func main() {
	for k, v := range os.Args[1:] {
		fmt.Printf("index: %d, value: %s\n", k, v)
	}
}

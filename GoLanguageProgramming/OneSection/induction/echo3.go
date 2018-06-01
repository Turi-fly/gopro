package main

import (
	"fmt"
	"os"
	"strings"
)

//不给GC带来很大的压力
func main() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}

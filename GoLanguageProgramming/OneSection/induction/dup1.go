package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	fmt.Print(">>> ")
	for input.Scan() {
		fmt.Print(">>> ")
		counts[input.Text()]++
		if input.Text() == "exit" {
			break
		}
	}
	//注意: 忽略 Input.err()中可能出现错误
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

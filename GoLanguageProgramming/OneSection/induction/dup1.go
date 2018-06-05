package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	`
	扫描器 bufio 里面的NewScanner扫描器从程序的标准输入,它可以以行或者单词为单位断开,这是以行为单位处理的最简单的方式.
	`
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

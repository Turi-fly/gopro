package main

import (
	"bufio"
	"fmt"
	"os"
)

func countLines(f *os.File, counts map[string]map[string]int) {
	a := make(map[string]int)
	input := bufio.NewScanner(f)
	for input.Scan() {
		a[input.Text()]++
	}
	counts[f.Name()] = a
	//注意:忽略input.Err()中的错误
}

func main() {
	counts := make(map[string]map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
		for filen, nmap := range counts {
			for line, n := range nmap {
				if n > 1 {
					fmt.Printf("%s\t%d\t%s\n", filen, n, line)
				}
			}
		}
	}

}

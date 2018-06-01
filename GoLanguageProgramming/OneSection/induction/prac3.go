package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	//高效
	s1 := time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "))
	s2 := time.Since(s1)
	fmt.Println("Program string time: ", s2)

	// 低效
	s3 := time.Now()
	var s, sep = "", ""
	for _, v := range os.Args[1:] {
		s += sep + v
		sep = " "
	}
	fmt.Println(s)
	s4 := time.Since(s3)
	fmt.Println("Program for time: ", s4)

}

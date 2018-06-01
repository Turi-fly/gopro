package main

import "fmt"

var v1 int
var v2 string
var v3 [10]int //数组
var v4 []int   //数组切片
var v5 struct {
	f int
}

var v6 *int           //指针
var v7 map[string]int //map
var v8 func(a int) int

func main() {
	fmt.Println(v1)
	fmt.Println(v2)
	fmt.Println(v3, len(v3), cap(v3))
	fmt.Println(v4, len(v4), cap(v4))
	fmt.Println(v6)
	fmt.Println(v7)
	fmt.Println(v8)
}

// result:
/*
0

[0 0 0 0 0 0 0 0 0 0] 10 10
[] 0 0
<nil>
map[]
<nil>
*/

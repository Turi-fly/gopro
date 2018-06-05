package main

import "fmt"

func GetName() (firstName, lastName, nickName string) {
	return "May", "chan", "Chibi Maruko"
}

func main() {
	//三种申明定义变量的方式
	var v1 int = 10
	var v2 = 10
	v3 := 10

	fmt.Println(v1, v2, v3)

	// var i int
	// i := 2
	//error:no new variables on left side of :=

	/////////////////////////////////////////
	//变量赋值和多变量赋值
	var v10 int
	v10 = 123

	fmt.Println(v10)

	// i, j = j, i
	// t = i; i = j ; j = t
	////////////////////////////////////////
	//匿名变量
	_, _, nickName := GetName()
	fmt.Println(nickName)

	//常量的定义
	const Pi float64 = 3.1415926 //有类型常量定义
	const zero = 0.0             //无类型浮点定义
	const (
		size int64 = 1024
		eof        = -1
	) //无类型整型定义
	const u, v float32 = 0, 3   //常量多赋值操作
	const a, b, c = 3, 4, "foo" // 无类型常量多赋值操作

	////////////////////////////////////////
	const ( // iota被重设为0
		c0 = iota // c0 == 0
		c1 = iota // c1 == 1
		c2 = iota // c2 == 2
	)

	const (
		a = 1 << iota // a == 1 (iota在每个const开头被重设为0)
		b = 1 << iota // b == 2
		c = 1 << iota // c == 4
	)

	const ()

}

package main

import (
	"fmt"
)

func f(p *int) int {
	*p++
	return *p
}

//非课程源码的 练习
func main() {
	//声明变量
	//var name type = expression
	// var s string = "test"
	// var y, x, z = "sz", 2, true
	// fmt.Println(s)
	// fmt.Printf("%s - %d - %v\n", y, x, z)
	v := 1
	f(&v)
	fmt.Println(f(&v))
}

package main

import (
	"fmt"
)

func f01(p *int) int {
	*p++
	return *p
}

//非课程源码的 练习
func main01() {
	//声明变量
	//var name type = expression
	var s string = "test"
	var y, x, z = "sz", 2, true
	fmt.Println(s)
	fmt.Printf("%s - %d - %v\n", y, x, z)
	v := 1
	f01(&v)
	fmt.Println(f01(&v))
}

// 2.3.3 newh函数
func newInt02() *int {
	//两种写法的有同样的行为
	return new(int)

	// var dummary int
	// return &dummary
}

func main02() {
	//new(T) 创建变量，地址类型为*T
	p := new(int)
	fmt.Println(p)
	// *p = 2
	// fmt.Println(*p)
	q := new(int)
	fmt.Println(q)
}

//2.3.4变量的生命周期
var global *int

func f03() {
	var x int
	x = 1
	global = &x
}

func g03() {
	y := new(int)
	*y = 1
}
func main03() {
	f03()
	fmt.Println(*global)
}

//2.4 赋值 计算斐波那契数列的第n个数
func fib04(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		fmt.Printf("前 i = %d, x = %d, y = %d\n", i, x, y)
		x, y = y, x+y
		fmt.Printf("后 i = %d, x = %d, y = %d\n", i, x, y)
	}
	return x
}

func main04() {
	p := fib04(7)
	fmt.Println("p = ", p)
}

//2.6.2 包的初始化
func f05(c int) int {
	return c + 1
}

func main05() {
	var c = 1
	var b = f05(c)
	var a = b + c
	fmt.Println(a)
}

//2.7 作用域，变量和函数不能同名
func f06() int {
	return 111
}

func main06() {
	s := "f"
	//f := "f" 这样写下面的f()就找不到了
	d := f06()
	fmt.Println(s)
	fmt.Println(d)
}

func main07() {
	// x := "hello!"
	// for i := 0; i < len(x); i++ {
	// 	x := x[i]
	// 	if x != '!' {
	// 		x := x + 'A' - 'a'
	// 		fmt.Printf("%c\n", x)
	// 	}
	// }

	// f := 3.141
	// i := int(f)
	// fmt.Println("i = ", i)
	// f = 1.99
	//fmt.Println(int(f))
	f := 1e100
	// i := int(f)
	fmt.Printf("%g\n", f)
}

//3.5 字符串
func main08() {
	s := "hello, world"
	fmt.Println(len(s))

	fmt.Println(s[0], s[7])
	fmt.Println(s[0:12])

	fmt.Println(s[:5])
	fmt.Println(s[7:])
	fmt.Printf("测试\a\n")
}

func main() {
	var a = "abc"
	var b = []byte(a)
	fmt.Println(b)
	s3 := string(b)
	fmt.Println(s3)
}

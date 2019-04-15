package main

import (
	"fmt"
	"os"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", comma(os.Args[i]))
	}
}

//递归函数 自己 调用 自己
//过程中打印的变量，有利于理解此程序
func comma(s string) string {
	n := len(s)
	fmt.Println("n = ", n)
	if n <= 3 {
		return s
	}
	fmt.Println("s = ", s)
	return comma(s[:n-3]) + "," + s[n-3:]
}

package main

import (
	"flag"
	"fmt"
	"strings"
)

// func Bool(name string, value bool, usage string) *bool
// Bool用指定的名称、默认值、使用信息注册一个bool类型flag。返回一个保存了该flag的值的指针。

// func String(name string, value string, usage string) *string
// String用指定的名称、默认值、使用信息注册一个string类型flag。返回一个保存了该flag的值的指针。

var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", " ", "separator")

func main() {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}

package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = ""
	}
	fmt.Println(s)
}

//练习作业 1.1-1.2
// func main() {
// 	s, sep := "", ""
// 	for i, arg := range os.Args[1:] {
// 		fmt.Printf("%d-%s\n", i, arg)
// 		s += sep + arg
// 		sep = ""
// 	}
// 	fmt.Println(s)
// }

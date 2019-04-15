package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		fmt.Println(basename(input.Text()))
	}
}

func basename(s string) string {
	// 将最后一个/ 和之前的部分全部舍弃
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}
	// 保留最后一个 点 之前的所有内容
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s
}

//basename2 的代码
//import strings

// func basename(s string) string {
// 	slash := strings.LastIndex(s, "/") // -1 if "/" not found
// 	s = s[slash+1:]
// 	if dot := strings.LastIndex(s, "."); dot >= 0 {
// 		s = s[:dot]
// 	}
// 	return s
// }

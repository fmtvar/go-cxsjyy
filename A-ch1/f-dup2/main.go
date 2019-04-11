//See page 8
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
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
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
	// NOTE: ignoring potential errors from input.Err()
}

//课后习题 1.4
// 这里引用一个结构体。来存储 文件名
// 然后把结构体放到map 中。line的内容做键值，结构体做值
// type line struct {
// 	FileName string
// 	String   string
// }

// func main() {
// 	counts := make(map[line]int)
// 	files := os.Args[1:]
// 	if len(files) == 0 {
// 		countLines(os.Stdin, counts, "ARGS")
// 	} else {
// 		for _, arg := range files {
// 			f, err := os.Open(arg)
// 			if err != nil {
// 				fmt.Fprintf(os.Stderr, "dup2: %vn", err)
// 				continue
// 			}
// 			countLines(f, counts, f.Name())
// 			f.Close()
// 		}
// 	}

// 	for line, n := range counts {
// 		if n > 1 {
// 			fmt.Printf("%st%st%dn", line.FileName, line.String, n)
// 		}
// 	}
// }

// func countLines(file *os.File, countsmap[line]int, fileNamestring) {
// 	input := bufio.NewScanner(file)
// 	for input.Scan() {
// 		counts[line{
// 			FileName: fileName,
// 			String:   input.Text(),
// 		}]++
// 	}
// }

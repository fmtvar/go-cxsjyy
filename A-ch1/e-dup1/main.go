//See page 6
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		//对16行代码解释
		//line := input.Text()
		// 这里是把读到的内容当做键，新内容，右边值就为1，
		//此时来一个重复内容。右边会在原来键值对值的基础上加1
		//因此可以标识重复的行
		//counts[line] = counts[line]+1
		counts[input.Text()]++

	}

	for line, n := range counts {
		if n > 1 {
			//\t制表符号  \n换行
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

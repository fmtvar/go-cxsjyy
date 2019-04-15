package main

import (
	"bytes"
	"fmt"
)

//如果要在 bytes.Buffer变量后面添加任意文字符号
//最好使用bytes.Buffer 的 WriteRune方法，而追加ASCII字符，如'[',']',则使用WriteByte亦可
func intsToString(values []int) string {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, v := range values {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']')
	return buf.String()
}

func main() {
	fmt.Println(intsToString([]int{1, 2, 3})) // "[1, 2, 3]"
}

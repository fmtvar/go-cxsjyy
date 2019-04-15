package main

import (
	"crypto/sha256"
	"fmt"
)

//prt 指向实现数组ss，它是指向数组，它是数组指针
//*p代表指针所指向的内存，就是实参a
func updDate(prt *[4]int) {
	for i, _ := range prt {
		prt[i] = 5555
	}
}

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)

	// 结果：
	// 2d711642b726b04401627ca9fbac32f5c8530fb1903cc4db02258717921a4881
	// 4b68ab3847feda7d6c62c1fbcbeebfa35eab7351ed5e78f4ddadea5df64b8015
	// false
	// [32]uint8

	//数组本来是值传递的。这里手动处理成 引用传递(指针传递)
	var ss = [4]int{1, 2, 3, 4}
	fmt.Println(ss)
	updDate(&ss)
	fmt.Println(ss)
}

package main

import (
	"fmt"
	//这是我自定义的路径，没按书上的来
	"go-cxsjyy/B-ch2/d-1-tempconv/utils"
	"os"
	"strconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		f := utils.Fahrenheit(t)
		c := utils.Celsius(t)
		fmt.Printf("%s = %s, %s = %s\n",
			f, utils.FToC(f), c, utils.CToF(c))
	}
}

package main

import (
	"fmt"
	"go-cxsjyy/B-ch2/d-0-tempconv/utils"
)

func main() {
	fmt.Printf("%g\n", utils.BoilingC-utils.FreezingC) // "100" °C
	boilingF := utils.CToF(utils.BoilingC)
	fmt.Printf("%g\n", boilingF-utils.CToF(utils.FreezingC)) // "180" °F

	c := utils.FToC(212.0)
	fmt.Println(c.String()) // "100°C"
	fmt.Printf("%v\n", c)   // "100°C"; no need to call String explicitly
	fmt.Printf("%s\n", c)   // "100°C"
	fmt.Println(c)          // "100°C"
	fmt.Printf("%g\n", c)   // "100"; does not call String
	fmt.Println(float64(c)) // "100"; does not call String
}

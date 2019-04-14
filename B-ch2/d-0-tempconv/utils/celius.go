package utils

import "fmt"

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

// 华氏温度  Fahrenheit(t) 是类型转换 不是函数调用
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// 摄氏温度  Celsius(t) 是类型转换 不是函数调用
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

// !-
func (c Celsius) String() string { return fmt.Sprintf("%g°C", c) }

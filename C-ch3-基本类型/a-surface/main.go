package main

import (
	"fmt"
	"math"
)

/**
 *说实话我没仔细看这个程序,
 * 等到这本书看第二遍或者第三遍的时候在啃这些吧。
 * 第一编先把 go 的基本知识掌握了
 */
const (
	width, height = 600, 320            // 以像素标示画布的大小
	cells         = 100                 // 网格单元的个数
	xyrange       = 30.0                // 坐标轴的范围 (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // x 或 y 轴上每个单位长度的像素
	zscale        = height * 0.4        // z 轴上每个单位长度的像素
	angle         = math.Pi / 6         // x、y轴的角度（=30）
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Println("</svg>")
}

func corner(i, j int) (float64, float64) {
	// 求出网格单元（i，j）的顶点坐标(x,y)
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// 计算曲面高度z
	z := f(x, y)

	// 将 (x,y,z) 等角投射到二维SVG绘图平面上，坐标是(sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // 到 (0,0) 的距离
	return math.Sin(r) / r
}

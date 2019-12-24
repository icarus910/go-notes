package main

import (
	"fmt"
	"image/color"
	"math/cmplx"
	"math/rand"
	lib "git.kuainiujinke.com/dsq/commonlib-go"
)

const MaxRand  = 100
type point struct {
	x, y int
}

func StatRandomNumbers(numRands int) (int, int) {
	fmt.Print("123")
	// 声明了两个变量（类型都为int，初始值都为0）
	var a, b int
	// 一个for循环代码块
	for i := 0; i < numRands; i++ {
		// 一个if-else条件控制代码块
		if rand.Intn(MaxRand) < MaxRand/2 {
			a = a + 1
		} else {
			b++ // 等价于：b = b + 1
		}
	}
	return a, b // 此函数返回两个结果
}

func main() {
	fmt.Println(lib.Now().Unix())
}



func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}
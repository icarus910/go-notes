package main

import (
	"fmt"
	"image/color"
	"math/cmplx"
	"math/rand"
)

const MaxRand = 100

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

func test() (a int32, b int32) {
	a = 1
	return

}

func main() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}

	LetCode(nums)
	fmt.Print(nums)
	a, b := 123, "Go"
	fmt.Printf("a == %v == 0x%x, b == %s\n", a, a, b)
	fmt.Printf("type of a: %T, type of b: %T\n", a, b)

}

func LetCode(nums []int) int {
	i, j := 0, 0
	length := len(nums)
	for j = 1; j < length; j ++ {
		if nums[j] != nums[i] {
			i ++
			nums[i] = nums[j]
		}
		fmt.Println(nums)
	}
	return i + 1
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

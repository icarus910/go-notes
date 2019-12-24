package main

// 声明了两个单独的有名常量。（是的，
// 非ASCII字符可以用做标识符。）
const π = 3.1416
const Pi = π // 等价于：Pi == 3.1416

// 声明了一组有名常量。
const (
	No         = !Yes
	Yes        = true
	MaxDegrees = 360
	Unit       = "弧度"
)

func main() {
	// 声明了三个局部有名常量。
	const DoublePi, HalfPi, Unit2 = π * 2, π * 0.5, "度"

	const k int16 = 255
	var n = k            //
	//var f = uint8(k + 1) // error 常量类型不可溢出
	var _ = uint8(n + 1) // 0 非常量可溢出
}

package constdemo

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
}

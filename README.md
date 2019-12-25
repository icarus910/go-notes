# go-notes

## 关键字
- const、func、import、package、type和var用来声明各种代码元素。
- chan、interface、map和struct用做 一些组合类型的字面表示中。
- break、case、continue、default、else、fallthrough、for、goto、if、range、return、select和switch用在流程控制语句中。
- defer和go也可以看作是流程控制关键字， 但它们有一些特殊的作用。

## 标识符
- **关键字不能被用做标识符。**
- _是一个特殊字符，它叫做空标识符

## 基本类型
- 布尔类型：bool。
- 11种内置整数类型：int8、uint8、int16、uint16、int32、uint32、int64、uint64、int、uint和uintptr。
- 两种内置浮点数类型：float32和float64。
- 两种内置复数类型：complex64和complex128。
- 一种内置字符串类型：string。
### 别名/特殊
- byte是uint8的内置别名。 我们可以将byte和uint8看作是同一个类型
- rune是int32的内置别名。 我们可以将rune和int32看作是同一个类型。
- iota是Go中预声明（内置）的一个特殊的有名常量。自增。 iota被预声明为0，但是它的值在编译阶段并非恒定。 当此预声明的iota出现在一个常量声明中的时候，它的值在第n个常量描述中的值为n（从0开始）。 所以iota只对含有多个常量描述的常量声明有意义。
#### _ 特殊标识符
- import 下划线：当导入一个包时，该包下的文件里所有init()函数都会被执行，然而，有些时候我们并不需要把整个包都导入进来，仅仅是是希望它执行init()函数而已。
- _下划线在代码中：忽略返回值（一般是错误值）；占位符；接口断言；函数中省略参数名；引入包执行init()；

````
string(65)          // "A"
string('A')         // "A"
string('\u68ee')    // "森"
string(-1)          // "\uFFFD"
string(0xFFFD)      // "\uFFFD"
string(0x2FFFFFFFF) // "\uFFFD"
```` 
## 常量变量
- 两个类型不一样的基本类型值是不能相互赋值的, 我们必须使用显式类型转换将一个值转换为另一个值的类型之后才能进行赋值。
````
const k int16 = 255
var n = k            //
var f = uint8(k + 1) // error 常量数字值的类型转换不能溢出
var h = uint8(n + 1) // 0 非常量可溢出
````

## 代码包和包引入
### fmt.Printf函数调用的输出格式
- %v：将被替换为对应实参字符串表示形式。
- %T：将替换为对应实参的类型的字符串表示形式。
- %x：将替换为对应实参的十六进制表示。实参的类型必须为整数，整数数组（array）或者整数切片（slice）等。
- %s：将被替换为对应实参的字符串表示形式。实参的类型必须为字符串或者字节切片（byte slice）类型。
- %%：将被替换为一个百分号。

## 基本流程控制
- switch-case：fallthrough关键字可以直接跳到下一个代码块
- 对于for/select/switch,Label必须紧挨着他们
````
func FindSmallestPrimeLargerThan(n int) int {
Outer:
	for n++; ; n++{
		for i := 2; ; i++ {
			switch {
			case i * i > n:
				break Outer //结束循环，跳到return n
			case n % i == 0:
				continue Outer//结束当前的执行，跳回for
			}
		}
	}
	return n
}
````

## Go的携程-goroutine
- 协程也被称为绿色线程。绿色线程是由程序的运行时（runtime）维护的线程。一个绿色线程的开销（比如内存消耗和情景转换）比一个系统线程常常小得多。
- Go不支持创建系统线程，所以协程是一个Go程序内部唯一的并发实现方式。
- 查询当前程序可利用的逻辑CPU数目:runtime.NumCPU()
- 获取和设置逻辑处理器的的数量:runtime.GOMAXPROCS()

### sync.WaitGroup
- Add方法用来注册新的需要完成的任务数。
- Done方法用来通知某个任务已经完成了。
- 一个Wait方法调用将阻塞（等待）到所有任务都已经完成之后才继续执行其后的语句。

### MPG模型
- M - worker thread, or machine. 工作线程 （线程）
- P - processor, 执行Go代码时所必须的一种资源。（逻辑处理）
- G - goroutinue. 协程

**P(逻辑处理器)** 会讲将不同的处于运行状态的**G(协程)**交给不同的**M(线程)**来执行

### 延迟函数defer
- 一个延迟调用可以修改包含此延迟调用的最内层函数的返回值
````
func Triple(n int) (r int) {
	defer func() {
		r += n // 修改返回值
		fmt.Println("1", r)
		fmt.Println("3", n)
	}()
	fmt.Println("2", n)
	return n + n // r = n + n
}

func main() {
	fmt.Println(Triple(5)) //15
}
````
- 一个匿名函数体内的表达式是在此函数被执行的时候才会被逐个估值的，不管此函数是被普通调用还是延迟/协程调用。
````
package main

import "fmt"

func main() {
	func() {
		for i := 0; i < 3; i++ {
			defer fmt.Println("a:", i)
		}
	}() ///210
	fmt.Println()
	func() {
		for i := 0; i < 3; i++ {
			defer func() {
				fmt.Println("b:", i)
			}()
		}
	}() //333
}
````

## 不同库print的区别
- log标准库中的打印函数是经过了同步处理，fmt标准库中的打印函数没有被同步



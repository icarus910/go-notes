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
- 下划线在代码中：忽略返回值（一般是错误值）；占位符；接口断言；函数中省略参数名；


````
string(65)          // "A"
string('A')         // "A"
string('\u68ee')    // "森"
string(-1)          // "\uFFFD"
string(0xFFFD)      // "\uFFFD"
string(0x2FFFFFFFF) // "\uFFFD"
````

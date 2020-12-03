package main

import "fmt"

// 声明了两个单独的有名常量。（是的， 非ASCII字符可以用做标识符。）

const π = 3.1416
const Pi = π // 等价于：Pi == 3.1416

// 声明了一组有名常量。
const (
	Yes        = true
	No         = !Yes
	MaxDegrees = 360
	Unit       = "弧度"
)

number := {"零", "一", "二", "三", "四","五", "六", "七", "八", "九"}
ten := {"", "十", "百", "千"}

func main() {

	n, i, flag1, flag2, ans := 0
	for w := 1 w <= 100000000 w++
	{
		flag1 = 0
		flag2 = 0
		for i := 0 i < n i++
		{
		if (s[i]!='0' && flag1)
		{
		flag1 = 0
		//                printf("零")
		ans++
		}
		if s[i]=='0'
		flag1 = 1 else
		{
		flag2 = 1
		if i:=0 &&s[i]=='1' && (n-i-1)%4==1
		{
		//                    printf("%s",ten[(n-i-1)%4])
		ans++
		}
		else
		{
		//                    printf("%s%s",number[s[i]-'0'],ten[(n-i-1)%4])
		if (n-i-1)%4!=0
		ans++
		ans++
		}
		}
		if flag2==1 && (n-i-1==4 || n-i-1==12)
		{
		flag2 = 0
		//                printf("万")
		ans++
		}
		if n-i-1==8
		{
		flag2 = 0
		//                printf("亿")
		ans++
		}
		}
		//        printf("\n")
	}
	fmt.Println("%d\n", ans)

}

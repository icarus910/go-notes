package main

import (
	"fmt"
	"time"
)
//通道元素值的传递都是复制过程
//对于官方编译器，最大元素尺寸为65535
//若通道发送数据协程队列和接收数据协程队列只要有一个不为空，则此通道不会被GC
func channel (){
	// 容量为0的通道为非缓冲通道，不为0则为缓冲通道
	ch := make(chan  int , 10)
	close(ch) //关闭channel
	var v int
	ch <- v // 发送值v
	v = <- ch // 接收值
	_ = cap(ch) // 查询容量
	_ = len(ch) // 查询长度（长度为当前已被发送到此通道但还未被接收出去的元素数量）
}

func demo() {
	c := make(chan int)
	go func(ch chan<- int,x int){
		time.Sleep(time.Second)
		ch <- x*x
	}(c,3)
	done := make(chan struct{})
	go func(ch <-chan int){
		n := <-ch
		fmt.Println(n)
		time.Sleep(time.Second)
		done <- struct{}{}
	}(c)
	<-done
	fmt.Println("end")
}

func football(){
	ball := make(chan string)
	kickBall := func(name string) {
		for{
			fmt.Println(<-ball,"传球","\n")
			time.Sleep(time.Second)
			ball <- name
		}
	}
	go kickBall("zzq")
	go kickBall("tge")
	go kickBall("gjw")
	ball <- "裁判"
	time.Sleep(time.Second*30)
	//var c chan bool
	//<-c
}

func fibonacci() {
	fibonacci := func() chan uint64 {
		c := make(chan uint64)
		go func() {
			var x, y uint64 = 0, 1
			for ; y < (1 << 63); c <- y { // 步尾语句
				x, y = y, x+y
			}
			close(c)
		}()
		return c
	}
	c := fibonacci()
	for x, ok := <-c; ok; x, ok = <-c { // 初始化和步尾语句
		time.Sleep(time.Second)
		fmt.Println(x)
	}
}

func main(){
	c := make(chan int,3)
	v,ok:= <- c
	fmt.Println(v,ok)
	close(c)
}
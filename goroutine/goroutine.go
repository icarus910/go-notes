package main

import (
	"fmt"
	"log"
	"runtime"
	"sync"
	"time"
)
var wg sync.WaitGroup

func SayGreetings(greeting string, times int) {
	//log.Println(greeting)
	//log.Println(runtime.NumGoroutine())
	//log.Println(time.Now())
	//for i := 0; i < times; i++ {
	//	log.Println(greeting)
	//	d := time.Second * time.Duration(rand.Intn(5)) / 2
	log.Println(greeting)
	time.Sleep(2*time.Second)
	//}
	wg.Done() // 通知当前任务已经完成。
}

func test() {
	var a = 123
	wg.Add(1)
	go func(x int) {
		time.Sleep(time.Second)
		fmt.Println(x, a) // 123 789
		wg.Done()
	}(a) 	//这里创建时a=x=123
	//速度过快a并不会重新赋值为789
	//time.Sleep(2 * time.Second) //123 123
	a = 789
	wg.Wait()

}

//MPG模型 M（线程） P（Processor）处理器
//一个main()会创建当前机器逻辑cpu数量的P和相关联的M
//新创建等待运行的G会唤醒一个P来进行处理，同时P会创建一个相关联的M（P将G分配给M）
//若M空闲，则也会像P一样放入空闲的M链表中
//可创建goroutine的数量只和内存有关每个goroutine大概占用8k～9k的内存
func main() {
	num := runtime.GOMAXPROCS(1)
	cpuNum := runtime.NumCPU()
	fmt.Println(num,cpuNum)
	//rand.Seed(time.Now().UnixNano())
	//log.SetFlags(0)
	wg.Add(11) // 注册新任务。
	go SayGreetings("1!", 10)
	go SayGreetings("2!", 10)
	go SayGreetings("3!", 10)
	go SayGreetings("4!", 10)
	go SayGreetings("5!", 10)
	go SayGreetings("6!", 10)
	go SayGreetings("7!", 10)
	go SayGreetings("8!", 10)
	go SayGreetings("9!", 10)
	go SayGreetings("10!", 10)
	go SayGreetings("11!", 10)
	wg.Wait() // 阻塞在这里，直到所有任务都已完成。
}
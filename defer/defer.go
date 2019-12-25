package main

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func Triple(n int) (r int) {
	defer func() {
		r += n // 修改返回值
		fmt.Println("1", r)
		fmt.Println("3", n)
	}()
	fmt.Println("2", n)
	return n + n // r = n + n
}

func niming() {
	func() {
		for i := 0; i < 3; i++ {
			defer fmt.Println("a:", i)
		}
	}()
	fmt.Println()
	func() {
		for i := 0; i < 3; i++ {
			defer func() {
				fmt.Println("b:", i)
			}()
		}
	}()
}

func test() {
	for i := 0; i < 3; i++ {
		fmt.Println("c:", i)
		go func() { //i为协程创建时的值，所以结果不定 ；若defer则先执行for和i++，所以结果为333
			fmt.Println("b:", i)
			wg.Done()
		}()
		//time.Sleep(2*time.Second)
	}

}

func main() {
	//func() {
	//for i := 0; i < 3; i++ {
	//	defer fmt.Println("a:", i)
	//}
	//	defer func() {i := 10 ;i+=1}()
	//}()
	//fmt.Println()
	rand.Seed(time.Now().UnixNano())
	log.SetFlags(0)
	wg.Add(3)
	test()
	wg.Wait()
}

package main

import (
	"fmt"
	"strconv"
)

func double(x *int) {
	*x += *x
	x = nil // 此行仅为讲解目的
}

func main() {
	var b int = 70
	fmt.Println(strconv.Itoa(b)) // 6
	a := 3
	double(&a)
	fmt.Println(a) // 6
	p := &a
	fmt.Println(a, p ,*p) // 12 false
	double(p)
	fmt.Println(a, p ,*p) // 12 false
}
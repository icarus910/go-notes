package main

import "fmt"

func main() {
	a := map[int][6]string{1: {"123456"},}
	var b [6]int
	fmt.Println(a)
	fmt.Println(len(a), cap(b))

	pm := &map[string]int{"C": 1972, "Go": 2009}
	ps := &[]string{"break", "continue"}
	pa := &[...]bool{false, true, true, false}
	fmt.Printf("%T\n", pm) // *map[string]int
	fmt.Printf("%T\n", ps) // *[]string
	fmt.Printf("%T\n", pa) // *[4]bool
}

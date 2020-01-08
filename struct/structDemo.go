package main

import "fmt"

type Book struct {
	title, author string
	pages         int
}

type Person struct {
	name string
	age         int
}
func main()  {
	book2 := Book{pages:3}

	book1 := Person{"abc",1}

	book1 = book2
	fmt.Println(book1,book2)
}
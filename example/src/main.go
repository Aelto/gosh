package main

import (
	"fmt"
)

@template t
func log(v t) {
	fmt.Println(v)
}

@template t y
func logTwo(v t, w y) {
	fmt.Println(v, w)
}

@template T
func doubleSlice(list T) {
	newList := T{}

	for _, el := range list {
		newList := append(newList, el * 2)
	}

	return newList
}

func main() {

	number := 5
	message := "Hello World!"

	log(number)@int
	log(number)@int
	log(message)@string
	logTwo(5, "Hello")@int string

	doubleSlice([]int{5, 6, 12})@[]int

	fmt.Println()

}
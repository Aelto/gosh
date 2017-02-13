package main

import (
	"fmt"
)

//@ template int
func goshlogint(v int) {
	fmt.Println(v)
}

//@ template string
func goshlogstring(v string) {
	fmt.Println(v)
}

//@ template int string
func goshlogTwointstring(v int, w string) {
	fmt.Println(v, w)
}

//@ template []int
func goshdoubleSlice_int(list []int) {
	newList := []int{}

	for _, el := range list {
		newList := append(newList, el * 2)
	}

	return newList
}

func main() {

	number := 5
	message := "Hello World!"

	goshlogint(number) // @int
	goshlogint(number) // @int
	goshlogstring(message) // @string
	goshlogTwointstring(5, "Hello") // @int string

	goshdoubleSlice_int([]int{5, 6, 12}) // @[]int

	fmt.Println()

}
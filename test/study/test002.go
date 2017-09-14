package main

import (
	"fmt"
)

var a int

func test() {
	a = 1
	fmt.Println(a)
}
func main() {
	test()
	a = 2
	fmt.Println(a)
}

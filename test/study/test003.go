package main

import (
	"fmt"
)

var (
	a int
	b int
	c int
)

func test() {
	a = 1
	b = 2
	c = 3
}

func main() {
	test()
	fmt.Println(a, b, c)
}

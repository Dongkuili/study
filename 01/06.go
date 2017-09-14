package main

import "fmt"

func main() {
	r := [...]int{99: -1}
	// 定义了一个含有100个元素的数组r，最后一个元素被初始化为-1，其它元素都是用0初始化
	fmt.Println(r)
}

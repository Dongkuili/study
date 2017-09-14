package main

import "fmt"

// 数组比较：包括元素类型、数组类型都可以进行比较
// 用==比较运算符来比较两个数组（!= 不想等运算符）
// 数组相等：只有当两个数组的所有元素都是相等的时候
func main() {
	a := [2]int{1, 2}
	b := [...]int{1, 2}
	c := [2]int{1, 3}
	fmt.Println(a == b, a == c, b == c) // "true false false"
	// 下面是错误操作
	d := [3]int{1, 2}
	fmt.Println(a == d) // compile error: cannot compare [2]int == [3]int
}

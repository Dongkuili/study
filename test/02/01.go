package main

import "fmt"

func main() {
	var arr2 [2][2]int // 表示一个二维静态数组
	arr2[1][1] = 2     // [1][1]=2 表示第2行、第2列的数字为2
	arr2[0][1] = 1     // [0][1]=1 表示第1行、第2列的数字为1
	// 其他没有定义的值默认为类型零值
	fmt.Println(arr2)
	// 经典的循环条件初始化/条件判断/循环后条件变化
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			fmt.Println(arr2[i][j])
		}
	}
}

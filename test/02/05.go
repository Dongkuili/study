package main

import "fmt"

// 也可以指定一个索引和对应值列表的方式进行初始化
func main() {
	type Currency int
	const (
		USD Currency = iota // 美元
		EUR                 // 欤元
		GBP                 // 英镑
		RMB                 // 人民币
	)
	symbol := [...]string{USD: "$", EUR: "€", GBP: "￡", RMB: "￥"}
	fmt.Println(USD, symbol[USD]) // "3 ￥"
}

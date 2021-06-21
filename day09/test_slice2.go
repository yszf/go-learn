package main

import "fmt"

func main() {
	var a []int
	a = append(a, 1) // 追加1个元素
	fmt.Println(a)
	a = append(a, 1, 2, 3) // 追加多个元素, 手写解包方式
	fmt.Println(a)
	a = append(a, []int{1, 2, 3}...) // 追加一个切片, 切片需要解包
	fmt.Println(a)
}

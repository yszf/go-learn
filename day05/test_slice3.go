package main

import "fmt"

func main() {
	s1 := []int{0, 1, 2, 3, 8: 100}   // 通过初始化表达式构造，可使用索引号
	fmt.Println(s1, len(s1), cap(s1)) // [0 1 2 3 0 0 0 0 100] 9 9

	s2 := make([]int, 6, 8)           // 使用make创建，指定len和cap值
	fmt.Println(s2, len(s2), cap(s2)) // [0 0 0 0 0 0] 6 8

	s3 := make([]int, 6)              // 省略cap，相当于cap=len
	fmt.Println(s3, len(s3), cap(s3)) // [0 0 0 0 0 0] 6 6
}

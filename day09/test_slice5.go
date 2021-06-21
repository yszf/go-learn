package main

import "fmt"

func main() {
	var a = []int{1, 2, 3}
	a = append(a, 0) // 切片扩展1个空间
	fmt.Println(a)
	copy(a[2+1:], a[2:]) // a[i:]向后移动1个位置
	a[2] = 5             // 设置新添加的元素
	fmt.Println(a)
}

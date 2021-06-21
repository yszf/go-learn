package main

import "fmt"

//用copy和append组合也可以实现在中间位置插入多个元素(也就是插入一个切片)
func main() {
	var a = []int{1, 2, 3}
	var x = []int{4, 5}
	a = append(a, x...)       // 为x切片扩展足够的空间
	fmt.Println(a)            // [1 2 3 4 5]
	copy(a[1+len(x):], a[1:]) // a[1:]向后移动len(x)个位置
	fmt.Println(a)            // [1 2 3 2 3]
	copy(a[1:], x)            // 复制新添加的切片
	fmt.Println(a)            // [1 4 5 2 3]
}

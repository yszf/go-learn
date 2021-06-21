package main

import "fmt"

// 在开头一般都会导致内存的重新分配，而且会导致已有的元素全部复制1次。
// 因此，从切片的开头添加元素的性能一般要比从尾部追加元素的性能差很多。

func main() {
	var a = []int{1, 2, 3}
	a = append([]int{0}, a...) // 在开头添加1个元素
	fmt.Println(a)
	a = append([]int{-3, -2, -1}, a...) // 在开头添加1个切片
	fmt.Println(a)
}

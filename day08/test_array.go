package main

import "fmt"

func main() {
	var a [3]int                    // 定义长度为3的int型数组, 元素全部为0
	var b = [...]int{1, 2, 3}       // 定义长度为3的int型数组, 元素为 1, 2, 3
	var c = [...]int{2: 3, 1: 2}    // 定义长度为3的int型数组, 元素为 0, 2, 3
	var d = [...]int{1, 2, 4: 5, 6} // 定义长度为6的int型数组, 元素为 1, 2, 0, 0, 5, 6

	fmt.Println(a) // 0 0 0
	fmt.Println(b) // 1 2 3
	fmt.Println(c) // 0 2 3
	fmt.Println(d) // 1 2 0 0 5 6
}

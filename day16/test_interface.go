package main

import "fmt"

func main() {
	var a = []interface{}{1, 2, 3}

	fmt.Println(a)
	fmt.Println(a...) // 当参数的可变参数是空接口类型时，传入空接口的切片时需要注意参数展开的问题

	var b = []int{1, 2, 3}
	fmt.Println(b)
//	fmt.Println(b...) // 无法将 'b' (类型 []int) 用作类型 []interface{}
}

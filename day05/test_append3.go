package main

import "fmt"

func main() {
	data := [...]int{0, 1, 2, 3, 4, 10: 0}
	s := data[:2:3]
	s = append(s, 100, 200)      // 一次 append 两个值，超出 s.cap 限制。
	fmt.Println(s, data)         // 重新分配底层数组，与原数组⽆无关。[0 1 100 200] [0 1 2 3 4 0 0 0 0 0 0]
	fmt.Println(&s[0], &data[0]) // 比对底层数组起始指针 0xc00000a300 0xc0000181e0
	fmt.Println(len(s), cap(s))  // 4 6
}

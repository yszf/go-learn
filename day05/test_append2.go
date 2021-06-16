package main

import "fmt"

func main() {
	data := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s := data[:3]
	s2 := append(s, 100, 200)                           // 添加多个值。
	fmt.Println(data)                                   // [0 1 2 100 200 5 6 7 8 9]
	fmt.Println(s)                                      // [0 1 2]
	fmt.Println(s2)                                     // [0 1 2 100 200]
	fmt.Printf("%p, %p, %p\n", &data, &s, &s2)          // 0xc00000e1e0, 0xc000004078, 0xc000004090
	fmt.Printf("%p, %p, %p\n", &data[0], &s[0], &s2[0]) // 0xc00000e1e0, 0xc00000e1e0, 0xc00000e1e0
}

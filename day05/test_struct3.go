package main

import "fmt"

func main() {
	type User struct {
		id   int
		name string
	}
	// ⽀支持 "=="、"!=" 相等操作符，可⽤用作 map 键类型
	m := map[User]int{
		User{1, "Tom"}: 100,
	}

	fmt.Println(m)
}

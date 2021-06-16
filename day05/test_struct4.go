package main

import "fmt"

func main() {
	// 可定义字段标签，⽤用反射读取。标签是类型的组成部分。
	var u1 struct {
		name string "username"
	}
	// var u2 struct{ name string }
	// u2 = u1 // Error: cannot use u1 (type struct { name string "username" }) as
	// // type struct { name string } in assignment
	u1.name = "hello"
	fmt.Println(u1)
}

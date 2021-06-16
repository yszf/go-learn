package main

import "fmt"

func main() {
	// 空结构 "节省" 内存，⽐比如⽤用来实现 set 数据结构，或者实现没有 "状态" 只有⽅方法的 "静态类"
	var null struct{}
	set := make(map[string]struct{})
	set["a"] = null
	fmt.Println(set) // map[a:{}]
}

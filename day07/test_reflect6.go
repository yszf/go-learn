package main

import (
	"fmt"
	"reflect"
)

// 可从基本类型获取所对应复合类型。 方法 Elem 可返回复合类型的基类型。
var (
	Int    = reflect.TypeOf(0)
	String = reflect.TypeOf("")
)

func main() {
	c := reflect.ChanOf(reflect.SendDir, String)
	fmt.Println(c) // chan<- string

	m := reflect.MapOf(String, Int)
	fmt.Println(m) // map[string]int

	s := reflect.SliceOf(Int)
	fmt.Println(s) // []int

	t := struct{ Name string }{}
	p := reflect.PtrTo(reflect.TypeOf(t))
	fmt.Println(p) // *struct { Name string }

	t2 := reflect.TypeOf(make(chan int)).Elem()
	fmt.Println(t2) // int
}

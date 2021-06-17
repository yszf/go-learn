package main

import (
	"fmt"
	"reflect"
)

// 接口是否为 nil，需要 tab 和 data 都为空。可使用 IsNil 方法判断 data 值

func main() {
	var p *int

	var x interface{} = p
	fmt.Println(x == nil) // false

	v := reflect.ValueOf(p)
	fmt.Println(v.Kind(), v.IsNil()) // ptr true
}

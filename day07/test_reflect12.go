package main

import (
	"fmt"
	"reflect"
)

// 除 struct，其他复合类型 array、slice、map 取值示例。

func main() {
	v := reflect.ValueOf([]int{1, 2, 3})
	for i, n := 0, v.Len(); i < n; i++ {
		fmt.Println(v.Index(i).Int()) // 1 2 3
	}
	fmt.Println("---------------------------")
	v = reflect.ValueOf(map[string]int{"a": 1, "b": 2})
	for _, k := range v.MapKeys() {
		fmt.Println(k.String(), v.MapIndex(k).Int()) // a 1 b 2
	}
}

package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Username string
	age      int
}

// 除具体的 Int、String 等转换方法，还可返回 interface{}。只是非导出字段无法使用，需
//用 CanInterface 判断一下
func main() {
	u := User{"Jack", 23}
	v := reflect.ValueOf(u)

	fmt.Println(v.FieldByName("Username").Interface()) // Jack
	fmt.Println(v.FieldByName("age").Interface())      // panic: reflect.Value.Interface: cannot return value obtained from unexported field or method
}

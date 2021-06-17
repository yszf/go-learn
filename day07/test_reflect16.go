package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

// 非导出字段无法直接修改，可改用指针操作。
type User struct {
	Username string
	age      int
}

func main() {
	u := User{"Jack", 23}

	p := reflect.ValueOf(&u).Elem()
	p.FieldByName("Username").SetString("Tom")

	f := p.FieldByName("age")
	fmt.Println(f.CanSet()) // false
	// 判断是否能获取地址。
	if f.CanAddr() {
		//	age := (*int)(unsafe.Pointer(f.UnsafeAddr()))
		age := (*int)(unsafe.Pointer(f.Addr().Pointer())) // 等同
		*age = 88
	}
	// 注意 p 是 Value 类型，需要还原成接⼝口才能转型。
	fmt.Println(u, p.Interface().(User)) // {Tom 88} {Tom 88}
}

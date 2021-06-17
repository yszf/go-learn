package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Username string
}

type Admin struct {
	User
	title string
}

// 如果是指针，应该先使用 Elem ⽅方法获取目标类型，指针本身是没有字段成员的

func main() {
	u := new(Admin)
	t := reflect.TypeOf(u)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	for i, n := 0, t.NumField(); i < n; i++ {
		f := t.Field(i)
		fmt.Println(f.Name, f.Type)
	}
}

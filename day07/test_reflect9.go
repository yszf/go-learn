package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Username string
	age      int
}

type Admin struct {
	User
	title string
}

// Value 和 Type 使⽤用方法类似，包括使用 Elem 获取指针目标对象。
func main() {
	u := &Admin{User{"Jack", 23}, "NT"}
	v := reflect.ValueOf(u).Elem()

	fmt.Println(v.FieldByName("title").String())   // 用转换⽅方法获取字段值 // NT
	fmt.Println(v.FieldByName("age").Int())        // 直接访问嵌入字段成员  // 23
	fmt.Println(v.FieldByIndex([]int{0, 1}).Int()) // 用多级序号访问嵌⼊入字段成员 // 23
}

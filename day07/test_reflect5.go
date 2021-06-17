package main

import (
	"fmt"
	"reflect"
)

// 字段标签可实现简单元数据编程
type User struct {
	Name string `field:"username" type:"nvarchar(20)"`
	Age  int    `field:"age" type:"tinyint"`
}

func main() {
	var u User
	t := reflect.TypeOf(u)
	f, _ := t.FieldByName("Name")

	fmt.Println(f.Tag)              // field:"username" type:"nvarchar(20)"
	fmt.Println(f.Tag.Get("field")) // username
	fmt.Println(f.Tag.Get("type"))  //
}

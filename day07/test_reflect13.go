package main

import (
	"fmt"
	"reflect"
)

// Value 某些方法没有遵循 "comma ok" 模式，而是返回 ZeroValue，因此需要用 IsValid 判断一下是否可用。

/*
func (v Value) FieldByName(name string) Value {
	v.mustBe(Struct)
	if f, ok := v.typ.FieldByName(name); ok {
		return v.FieldByIndex(f.Index)
	}
	return Value{}
}
*/

type User struct {
	Username string
	age      int
}

func main() {
	u := User{}
	v := reflect.ValueOf(u)
	f := v.FieldByName("a")
	fmt.Println(f.Kind(), f.IsValid()) // invalid false
}

package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Username string
	age      int
}

/*
将对象转换为接口，会发生复制行为。该复制品只读，无法被修改。所以要通过接口改变
目标对象状态，必须是 pointer-interface。
就算是指针，我们依然没法将这个存储在 data 的指针指向其他对象，只能透过它修改目
标对象。因为目标对象并没有被复制，被复制的只是指针。
*/

func main() {
	u := User{"Jack", 23}

	v := reflect.ValueOf(u)
	p := reflect.ValueOf(&u)

	fmt.Println(v.CanSet(), v.FieldByName("Username").CanSet()) // false false

	fmt.Println(p.CanSet(), p.Elem().FieldByName("Username").CanSet()) // false true
}

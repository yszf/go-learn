package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Username string
	age      int
}

func main() {
	u := User{"Jack", 23}
	v := reflect.ValueOf(u)
	f := v.FieldByName("age")
	if f.CanInterface() { // false
		fmt.Println(f.Interface())
	} else {
		fmt.Println(f.Int()) // 23
	}
}

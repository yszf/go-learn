package main

import (
	"fmt"
	"reflect"
)

type User struct {
}
type Admin struct {
	User
}

func (*User) ToString() {}

func (Admin) Test() {}

// value-interface 和 pointer-interface 也会导致方法集存在差异。
func main() {
	var u Admin
	methods := func(t reflect.Type) {
		for i, n := 0, t.NumMethod(); i < n; i++ {
			m := t.Method(i)
			fmt.Println(m.Name)
		}
	}
	fmt.Println("--- value interface ---")
	methods(reflect.TypeOf(u)) // Test
	fmt.Println("--- pointer interface ---")
	methods(reflect.TypeOf(&u)) // Test ToString
}

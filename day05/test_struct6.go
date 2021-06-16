package main

import "fmt"

func main() {
	// 匿名字段
	// 匿名字段不过是⼀一种语法糖，从根本上说，就是⼀一个与成员类型同名 (不含包名) 的字段。 被匿名嵌⼊入的可以是任何类型，当然也包括指针。
	type User struct {
		name string
	}

	type Manager struct {
		User
		title string
	}

	m := Manager{
		User:  User{"Tom"}, // 匿名字段的显式字段名，和类型名相同。
		title: "Administrator",
	}

	fmt.Println(m) // {{Tom} Administrator}
}

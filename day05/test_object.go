package main

import (
	"fmt"
)

func main() {
	// ⾯面向对象三大特征里，Go 仅支持封装，尽管匿名字段的内存布局和行为类似继承。没有 class 关键字，没有继承、多态等等。
	type User struct {
		id   int
		name string
	}

	type Manager struct {
		User
		title string
	}

	m := Manager{User{1, "Tom"}, "Administrator"}
	// var u User = m // Error: cannot use m (type Manager) as type User in assignment
	// 没有继承，自然也不会有多态。
	var u User = m.User // 同类型拷贝。

	fmt.Println(m) // {{1 Tom} Administrator}
	fmt.Println(u) // {1 Tom}

	// 内存布局和 C struct 相同，没有任何附加的 object 信息
}

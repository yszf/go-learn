package main

import "fmt"

type User struct {
	id   int
	name string
}

type Manager struct {
	User
	title string
}

func (self *Manager) ToString() string {
	return fmt.Sprintf("Manager: %p, %v", self, self)
}

// 通过匿名字段，可获得和继承类似的复⽤用能力。依据编译器查找次序，只需在外层定义同名方法，就可以实现 "override"
func (self *User) ToString() string {
	return fmt.Sprintf("User: %p, %v", self, self)
}

func main() {
	m := Manager{User{1, "Tom"}, "Administrator"}
	fmt.Println(m.ToString()) // Manager: 0xc0000763c0, &{{1 Tom} Administrator}

	fmt.Println(m.User.ToString()) // User: 0xc0000763c0, &{1 Tom}
}

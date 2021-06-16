package main

import "fmt"

// 接口是一个或多个方法签名的集合，任何类型的方法集中只要拥有与之对应的全部方法，
// 就表示它 "实现" 了该接口，无须在该类型上显式添加接口声明。

// 所谓对应方法，是指有相同名称、参数列表 (不包括参数名) 以及返回值。当然，该类型还可以有其他方法
type Stringer interface {
	String() string
}

type Printer interface {
	Stringer // 接口嵌入
	Print()
}

type User struct {
	id   int
	name string
}

func (self *User) String() string {
	return fmt.Sprintf("user %d, %s", self.id, self.name)
}

func (self *User) Print() {
	fmt.Println(self.String())
}

func main() {
	var t Printer = &User{1, "Tom"} // *User方法集包含String、Print
	t.Print()
}

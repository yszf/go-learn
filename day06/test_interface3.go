package main

import "fmt"

// 匿名接⼝口可⽤用作变量类型，或结构成员
type Tester struct {
	s interface {
		String() string
	}
}

type User struct {
	id   int
	name string
}

func (self *User) String() string {
	return fmt.Sprintf("user %d, %s", self.id, self.name)
}

func main() {
	t := Tester{&User{1, "Tom"}}
	fmt.Println(t.s.String()) // user 1, Tom
}

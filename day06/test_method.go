package main

type Queue struct {
	elements []interface{}
}

// ⽅方法总是绑定对象实例，并隐式将实例作为第一实参 (receiver)。
func NewQueue() *Queue { // 创建对象实例
	return &Queue{make([]interface{}, 10)}
}

func (*Queue) Push(e interface{}) error { // 省略reciver 参数名
	panic("not implemented")
}

// func (Queue) Push(e int) error { // Error: method Push already declared for type
// 	panic("not implemented")
// }

func (self *Queue) length() int { // receiver 参数名可以是 self、this或其他
	return len(self.elements)
}

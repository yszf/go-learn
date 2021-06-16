package main

type Tester interface {
	Do()
}

// 让函数直接 "实现" 接口能省不少事。

type FuncDo func()

func (self FuncDo) Do() { self() }

func main() {
	var t Tester = FuncDo(func() { println("Hello, World!") })
	t.Do()
}

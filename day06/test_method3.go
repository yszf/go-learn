package main

// 从 1.4 开始，不再⽀支持多级指针查找⽅方法成员。

type X struct{}

func (*X) test() {
	println("X.test")
}

func main() {
	p := &X{}
	p.test()
	// Error: calling method with receiver &p (type **X) requires explicit dereference
	// (&p).test()
}

package main

func test() {
	defer func() {
		if err := recover(); err != nil {
			println(err.(string)) // 将 interface{} 转型为具体类型。 // panic error!
		}
	}()
	panic("panic error!")
}

func main() {
	test()
}

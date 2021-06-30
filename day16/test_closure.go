package main

func main() {
	for i := 0; i < 5; i++ {
		defer func() {
			println(i) // 5 5 5 5 5 闭包错误引用同一个变量
		}()
	}
}

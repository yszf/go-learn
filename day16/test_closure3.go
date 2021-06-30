package main

func main() {
	for i := 0; i < 5; i++ {
		defer func(i int) {
			println(i) // 通过函数参数传入 4 3 2 1 0
		}(i)
	}
}

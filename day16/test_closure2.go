package main

func main() {
	for i := 0; i < 5; i++ {
		i := i
		defer func() {
			println(i) // 改进的方法是在每轮迭代中生成一个局部变量 4 3 2 1 0
		}()
	}
}

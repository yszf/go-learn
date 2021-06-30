package main

func main() {
	m := map[string]string{
		"1": "1",
		"2": "2",
		"3": "3",
	}

	for k, v := range m {
		println(k, v) // map是一种hash表实现，每次遍历的顺序都可能不一样。
	}
}
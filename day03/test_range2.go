package main

func main() {
	s := []int{1, 2, 3, 4, 5}
	for i, v := range s { // 复制 struct slice { pointer, len, cap }。
		if i == 0 {
			s = s[:3]  // 对 slice 的修改，不会影响 range。
			s[2] = 100 // 对底层数据的修改。
		}
		println(i, v)
	}

	println("------------------------")
	for i, v := range s {
		println(i, v)
	}
	println("------------------------")

	s = []int{1, 2, 3, 4, 5}
	l := len(s)
	for i := 0; i < l; i++ {
		if i == 0 {
			s = s[:3]
			s[2] = 200
		}
		println(i, s[i]) // panic: runtime error: index out of range [3] with length 3
	}
}

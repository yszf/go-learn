package main

import "fmt"

// 将多个append操作组合起来，实现在切片中间插入元素

func main() {
	var a = []int{1, 2, 3}
	// 每个添加操作中的第二个append调用都会创建一个临时切片，
	// 并将a[i:]的内容复制到新创建的切片中，然后将临时创建的切片再追加到a[:i]。
	a = append(a[:1], append([]int{0}, a[1:]...)...) // 在第1个位置插入0
	fmt.Println(a)
	a = append(a[:1], append([]int{1, 2, 3}, a[1:]...)...) // 在第1个位置插入切片
	fmt.Println(a)
}

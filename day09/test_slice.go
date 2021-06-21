package main

import "fmt"

type SliceHeader struct {
	Data uintptr
	Len  int
	Cap  int
}

func main() {

	var (
		a []int               // nil切片, 和 nil 相等, 一般用来表示一个不存在的切片
		b = []int{}           // 空切片, 和 nil 不相等, 一般用来表示一个空的集合
		c = []int{1, 2, 3}    // 有3个元素的切片, len和cap都为3
		d = c[:2]             // 有2个元素的切片, len为2, cap为3
		e = c[0:2:cap(c)]     // 有2个元素的切片, len为2, cap为3
		f = c[:0]             // 有0个元素的切片, len为0, cap为3
		g = make([]int, 3)    // 有3个元素的切片, len和cap都为3
		h = make([]int, 2, 3) // 有2个元素的切片, len为2, cap为3
		i = make([]int, 0, 3) // 有0个元素的切片, len为0, cap为3
	)

	fmt.Println(a, a == nil) // [] true
	fmt.Println(b, b == nil) // [] false
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f, f == nil) // [] false
	fmt.Println(g)
	fmt.Println(h)
	fmt.Println(i, i == nil) // false

}

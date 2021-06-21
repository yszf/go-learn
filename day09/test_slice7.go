package main

import "fmt"

func main() {
	var a = []int{1, 2, 3, 4, 5}
	var b = a
	b = b[:len(b)-1] // 删除尾部1个元素
	fmt.Println(b)
	b = b[:len(b)-2] // 删除尾部2个元素
	fmt.Println(b)
	fmt.Println("----------------------------")

	var c = a
	c = c[1:] // 删除头部1个元素
	fmt.Println(c)
	c = c[2:] // 删除头部2个元素
	fmt.Println(c)
	fmt.Println("----------------------------")

	// 删除开头的元素也可以不移动数据指针，但是将后面的数据向开头移动。
	// 可以用append原地完成（所谓原地完成是指在原有的切片数据对应的内存区间内完成，不会导致内存空间结构的变化）
	var d = a
	fmt.Println(d)
	d = append(d[:0], d[1:]...) // 删除开头1个元素
	fmt.Println(d)
	d = append(d[:0], d[2:]...) // 删除开头2个元素
	fmt.Println(d)
	fmt.Println("----------------------------")

	// 可以用copy完成删除开头的元素
	var e = []int{1, 2, 3, 4, 5}
	fmt.Println(e)
	e = e[:copy(e, e[1:])] // 删除开头1个元素
	fmt.Println(e)
	e = e[:copy(e, e[2:])] // 删除开头2个元素
	fmt.Println(e)
	fmt.Println("----------------------------")

	// 对于删除中间的元素，需要对剩余的元素进行一次整体挪动，同样可以用append或copy原地完成：
	var f = []int{1, 2, 3, 4, 5}
	fmt.Println(f)
	f = append(f[:2], f[2+1:]...) // 删除中间1个元素
	fmt.Println(f)
	f = append(f[:2], f[2+2:]...) // 删除中间2个元素
	fmt.Println(f)
	fmt.Println("----------------------------")

	var g = []int{1, 2, 3, 4, 5}
	fmt.Println(g)
	var n1 = copy(g[2:], g[2+1:])
	fmt.Println("n1 = ", n1)
	g = g[:2+n1] // 删除中间1个元素
	fmt.Println(g)
	var n2 = copy(g[2:], g[2+2:])
	fmt.Println("n2 = ", n2)
	g = g[:2+n2] // 删除中间2个元素
	fmt.Println(g)

	// 删除开头的元素和删除尾部的元素都可以认为是删除中间元素操作的特殊情况。
}

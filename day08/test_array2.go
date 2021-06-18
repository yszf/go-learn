package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"
)

func main() {
	var a = [...]int{1, 2, 3} // a 是一个数组
	var b = &a                // b 是指向数组的指针

	fmt.Println(a[0], a[1]) // 打印数组的前2个元素
	fmt.Println(b[0], b[1]) // 通过数组指针访问数组元素的方式和数组类似

	for i, v := range b { // 通过数组指针迭代数组的元素
		fmt.Println(i, v)
	}

	var times [5][0]int
	for range times {
		fmt.Println("hello")
	}

	// 字符串数组
	var s1 = [2]string{"hello", "world"}
	var s2 = [...]string{"你好", "世界"}
	var s3 = [...]string{1: "世界", 0: "你好"}
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)

	// 结构体数组
	var line1 [2]image.Point
	var line2 = [...]image.Point{image.Point{X: 0, Y: 0}, image.Point{X: 1, Y: 1}}
	var line3 = [...]image.Point{{0, 0}, {1, 1}}
	fmt.Println(line1)
	fmt.Println(line2)
	fmt.Println(line3)

	// 图像解码器数组
	var decoder1 [2]func(io.Reader) (image.Image, error)
	var decoder2 = [...]func(io.Reader) (image.Image, error){
		png.Decode,
		jpeg.Decode,
	}
	fmt.Println(decoder1)
	fmt.Println(decoder2)

	// 接口数组
	var unknown1 [2]interface{}
	var unknown2 = [...]interface{}{123, "你好"}
	fmt.Println(unknown1)
	fmt.Println(unknown2)

	// 管道数组
	var chanList = [2]chan int{}
	fmt.Println(chanList)

	fmt.Printf("a: %T\n", a)  // a: [3]int
	fmt.Printf("a: %#v\n", a) // a: [3]int{1, 2, 3}

	fmt.Printf("b: %T\n", b)  // b: *[3]int
	fmt.Printf("b: %#v\n", b) // b: &[3]int{1, 2, 3}
}

package main

import (
	"fmt"
	"reflect"
)

// 可获取方法参数、返回值类型等信息。
type Data struct {
}

func (*Data) Test(x, y int) (int, int) {
	return x + 100, y + 100
}

func (*Data) Sum(s string, x ...int) string {
	c := 0
	for _, n := range x {
		c += n
	}
	return fmt.Sprintf(s, c)
}

func info(m reflect.Method) {
	t := m.Type
	fmt.Println(m.Name) // 方法名
	for i, n := 0, t.NumIn(); i < n; i++ {
		fmt.Printf(" in[%d] %v\n", i, t.In(i)) // 参数
	}
	for i, n := 0, t.NumOut(); i < n; i++ {
		fmt.Printf(" out[%d] %v\n", i, t.Out(i)) // 返回值
	}
}

func main() {
	d := new(Data)
	t := reflect.TypeOf(d)
	test, _ := t.MethodByName("Test")
	info(test)
	sum, _ := t.MethodByName("Sum")
	info(sum)
}

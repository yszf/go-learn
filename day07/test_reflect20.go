package main

import (
	"fmt"
	"reflect"
)

type Data struct {
}

func (*Data) Test(x, y int) (int, int) {
	return x + 100, y + 100
}

// 非导出方法无法调用，甚至无法透过指针操作，因为接口类型信息中没有该方法地址。
func (*Data) Sum(s string, x ...int) string { // sum (panic: reflect: call of reflect.Value.CallSlice on zero Value)
	c := 0
	for _, n := range x {
		c += n
	}
	return fmt.Sprintf(s, c)
}

// 如改用 CallSlice，只需将变参打包成 slice 即可。
func main() {
	d := new(Data)
	v := reflect.ValueOf(d)
	m := v.MethodByName("Sum")
	in := []reflect.Value{
		reflect.ValueOf("result = %d"),
		reflect.ValueOf([]int{1, 2}), // 将变参打包成 slice。
	}
	out := m.CallSlice(in)
	for _, v := range out {
		fmt.Println(v.Interface())
	}
}

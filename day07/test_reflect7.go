package main

import (
	"fmt"
	"reflect"
)

type Data struct {
}

func (*Data) String() string {
	return ""
}

// 方法 Implements 判断是否实现了某个具体接口，AssignableTo、ConvertibleTo 用于赋值和转换判断
func main() {
	var d *Data
	t := reflect.TypeOf(d)

	// 没法直接获取接⼝口类型，好在接⼝口本⾝身是个 struct，创建
	// ⼀一个空指针对象，这样传递给 TypeOf 转换成 interface{}
	// 时就有类型信息了。。
	it := reflect.TypeOf((*fmt.Stringer)(nil)).Elem()

	// 为啥不是 t.Implements(fmt.Stringer)，完全可以由编译器⽣生成。
	fmt.Println(t.Implements(it)) // true
}

package main

import (
	"fmt"
	"reflect"
)

type Data struct {
	b byte
	x int32
}

// 某些时候，获取对齐信息对于内存自动分析是很有用的
func main() {
	var d Data
	t := reflect.TypeOf(d)

	fmt.Println(t.Size(), t.Align()) // sizeof，以及最宽字段的对齐模数。 // 8 4

	f, _ := t.FieldByName("b")
	fmt.Println(f.Type.FieldAlign()) // 字段对齐。 // 1
}

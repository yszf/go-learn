package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Username string
}

type Admin struct {
	User
	title string
}

// 没有运行期类型对象，实例也没有附加字段用来表明身份。只有转换成接口时，才会在其
// itab 内部存储与该类型有关的信息，Reflect 所有操作都依赖于此。

// 以 struct 为例，可获取其全部成员字段信息，包括非导出和匿名字段。

func main() {
	var u Admin
	t := reflect.TypeOf(u)
	for i, n := 0, t.NumField(); i < n; i++ {
		f := t.Field(i)
		fmt.Println(f.Name, f.Type)
		if "User" == f.Name {
			for j, n2 := 0, f.Type.NumField(); j < n2; j++ {
				f2 := f.Type.Field(j)
				fmt.Println(f2.Name, f2.Type)
			}
		}
	}
}

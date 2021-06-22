package main

import (
	"errors"
	"fmt"
)

func foo() (err error) {
	defer func() {
		if r := recover(); r != nil {
			switch x := r.(type) {
			case string:
				err = errors.New(x)
			case error:
				err = x
			default:
				err = fmt.Errorf("Unknown panic: %v", r)
			}
		}
	}()

	panic("TODO")
}

func main() {
	defer func() {
		if r := recover(); r != nil {

		}
		// 虽然总是返回nil, 但是可以恢复异常状态
	}()

	foo()

	// 警告: 用`nil`为参数抛出异常
	panic(nil)

}

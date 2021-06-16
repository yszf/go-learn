package main

import "fmt"

func main() {
	var pass string
	fmt.Println("请输入密码：")
	fmt.Scan(&pass)
	if "111111" == pass {
		fmt.Println("请再次输入密码：")
		fmt.Scan(&pass)
		if "111111" == pass {
			fmt.Println("输入密码正确")
		} else {
			fmt.Println("密码错误")
		}
	} else {
		fmt.Println("密码错误")
	}
}

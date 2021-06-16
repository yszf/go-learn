package main

import "fmt"

func main() {
	s := "abcd"
	bs := []byte(s)
	bs[1] = 'B'
	println(string(bs))

	u := "电脑"
	us := []rune(u)

	us[1] = '话'
	println(string(us))

	s = "abc汉字"
	for i := 0; i < len(s); i++ { // byte
		fmt.Printf("%c,", s[i])
	}
	fmt.Println()
	for _, r := range s { // rune
		fmt.Printf("%c,", r)
	}
}

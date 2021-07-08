package main

import (
	"fmt"
	"github.com/bitly/go-simplejson"
)

func main() {

	js, _ := simplejson.NewJson([]byte(`{
		"test": {
			"array": [1, "2", 3],
			"int": 10,
			"float": 5.150,
			"bignum": 9223372036854775807,
			"string": "simplejson",
			"bool": true
		}
	}`))

	arr, _ := js.Get("test").Get("array").Array()
	fmt.Println(arr) // [1 2 3]

	i, _ := js.Get("test").Get("int").Int()
	fmt.Println(i) // 10

	fl, _ := js.Get("test").Get("float").Float64()
	fmt.Println(fl) // 5.15

	big, _ := js.Get("test").Get("bignum").Int64()
	fmt.Println(big) // 9223372036854775807

	ms := js.Get("test").Get("string").MustString()
	fmt.Println(ms) // simplejson

	bo, _ := js.Get("test").Get("bool").Bool()
	fmt.Println(bo) // true

}

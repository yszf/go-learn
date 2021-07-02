package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func sayHelloName3(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm() //解析url传递的参数，对于POST则解析响应包的主体（request body）
	if err != nil {
		log.Fatal(err)
	}
	//注意:如果没有调用ParseForm方法，下面无法获取表单的数据
	fmt.Println(r.Form) //这些信息是输出到服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	_, err = fmt.Fprintf(w, "Hello astaxie!") //这个写入到w的是输出到客户端的
	if err != nil {
		log.Fatal(err)
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //获取请求的方法
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.gtpl")
		log.Println(t.Execute(w, nil))
	} else {
		r.ParseForm()
		//请求的是登录数据，那么执行登录的逻辑判断
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])

		fmt.Fprintf(w, "username = %s\n", r.FormValue("username"))
		fmt.Fprintf(w, "password = %s\n", r.FormValue("password"))

		v := url.Values{}
		v.Set("name", "Ava")
		v.Add("friend", "Jess")
		v.Add("friend", "Sarah")
		v.Add("friend", "Zoe")
		// v.Encode() == "name=Ava&friend=Jess&friend=Sarah&friend=Zoe"
		fmt.Println(v.Get("name"))   // Ava
		fmt.Println(v.Get("friend")) // Jess
		fmt.Println(v["friend"])     // [Jess Sarah Zoe]
	}
}

func main() {
	http.HandleFunc("/", sayHelloName3)      //设置访问的路由
	http.HandleFunc("/login", login)         //设置访问的路由
	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

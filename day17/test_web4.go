package main

import (
	"crypto/md5"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"
)

func login2(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //获取请求的方法
	if r.Method == "GET" {
		timestamp := strconv.Itoa(time.Now().Nanosecond())
		hashWr := md5.New()
		hashWr.Write([]byte(timestamp))
		token := fmt.Sprintf("%x", hashWr.Sum(nil))

		t, _ := template.ParseFiles("login.gtpl")
		t.Execute(w, token)
	} else {
		//请求的是登陆数据，那么执行登陆的逻辑判断
		r.ParseForm()
		token := r.Form.Get("token") // 防止表单重复提交，解决方案是在表单中添加一个带有唯一值的隐藏字段。
		// 在验证表单时，先检查带有该唯一值的表单是否已经递交过了。
		// 如果是，拒绝再次递交；如果不是，则处理表单进行逻辑处理。
		if token != "" {
			//验证token的合法性
		} else {
			//不存在token报错
		}
		fmt.Println("token:", token)
		fmt.Println("username length:", len(r.Form["username"][0]))
		fmt.Println("username:", template.HTMLEscapeString(r.Form.Get("username"))) //输出到服务器端
		fmt.Println("password:", template.HTMLEscapeString(r.Form.Get("password")))
		template.HTMLEscape(w, []byte(r.Form.Get("username"))) //输出到客户端
	}
}

func main() {
	http.HandleFunc("/login", login2)        //设置访问的路由
	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

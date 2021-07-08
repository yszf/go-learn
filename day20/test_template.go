package main

import (
	"html/template"
	"os"
)

type Person struct {
	UserName string
	email    string //未导出的字段，首字母是小写的
}

func main() {
	t := template.New("fieldname example")
	//	t, _ = t.Parse("hello {{.UserName}}!")
	t, _ = t.Parse("hello {{.UserName}}! {{.}}  {{.email}} {{.}} ")
	p := Person{UserName: "Astaxie", email: "abc@163.com"}
	t.Execute(os.Stdout, p) // hello Astaxie! {Astaxie abc@163.com}
}

/*
使用Parse代替ParseFiles，因为Parse可以直接测试一个字符串，而不需要额外的文件
不使用handler来写演示代码，而是每个测试一个main，方便测试
使用os.Stdout代替http.ResponseWriter，因为os.Stdout实现了io.Writer接口

如果我们调用了一个不存在的字段是不会报错的，而是输出为空。
如果模板中输出{{.}}，这个一般应用于字符串对象，默认会调用fmt包输出字符串的内容。
*/

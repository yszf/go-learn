package main

import (
	"fmt"
	"log"
	"net/http"
)

type MyMux struct {
}

func (p *MyMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		sayHelloName2(w, r)
		return
	}
	http.NotFound(w, r)
	return
}

func sayHelloName2(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm() //解析参数，默认是不会解析的
	if err != nil {
		log.Fatal(err)
	}
	_, err = fmt.Fprintf(w, "Hello myroute!")
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	mux := &MyMux{}
	err := http.ListenAndServe(":9090", mux)
	if err != nil {
		log.Fatal(err)
	}
}

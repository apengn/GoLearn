package main

import (
	"net/http"
	"fmt"
	"log"
	"time"
	"io"
)

type MyHander struct {
}

//handle默认执行的方法
func (p *MyHander) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//根据map转发路由
	if v, ok := mux[r.URL.String()]; ok {
		v(w, r)
		return
	}
	io.WriteString(w, "hello"+r.URL.Path)
}
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello world Home"+r.URL.Path)
}
func User(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello world User"+r.URL.Path)
}

//定义全局变量
var mux map[string]func(http.ResponseWriter, *http.Request)

func main() {
	server := http.Server{Addr: ":8080",
		Handler: &MyHander{},
		ReadTimeout: 2 * time.Second}
	mux = make(map[string]func(http.ResponseWriter, *http.Request))

	mux["/home"] = Home
	mux["/user"] = User

	err := server.ListenAndServe()

	//mux := http.NewServeMux()
	//mux.Handle("/", &MyHander{})
	//mux.HandleFunc("/home", Home)
	//err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}

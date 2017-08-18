package main

import (
	"net/http"
	"fmt"
)

type MyMux struct{

}
type MyMuxs interface{


}
func sayHelloWorldWeb1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "HelloWorldWeb----Go") //返回响应输出打客户端的
}
func (p *MyMux) ServeHttp(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path == "/" {
		sayHelloWorldWeb1(w, r)
		return
	}
	http.NotFound(w, r)
	return
}
func main() {
	http.ListenAndServe(":90910",nil)
}

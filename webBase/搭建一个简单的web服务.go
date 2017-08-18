package main

//导入使用的依赖包
import (
	"net/http"
	"fmt"
	"log"
)

func sayHelloWorldWeb(w http.ResponseWriter,r  *http.Request)  {
	r.ParseForm()//解析表单 参数  默认是不会解析的
	fmt.Println(r.Form)//表单数据
	fmt.Println("path:",r.URL.Path)
	fmt.Println("scheme:",r.URL.Scheme)
	fmt.Println("Host:",r.URL.Host)
	fmt.Println("Port",r.URL.Port())

	for k,v:=range r.Form{
		fmt.Println("key",k)
		fmt.Println("values:",v)
	}
	fmt.Fprintf(w,"HelloWorldWeb----Go")//返回响应输出打客户端的

}

func main() {
	http.HandleFunc("/",sayHelloWorldWeb)//设置访问的路由
	err:=http.ListenAndServe(":91090",nil)//设置监听的的端口
	if err!=nil {
		log.Fatal(err)
	}
}


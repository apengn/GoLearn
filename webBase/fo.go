package main

import (
	"net/http"
	"fmt"
	"html/template"
	"log"
)
//验证表单的输入
func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request method", r.Method)
	if r.Method == "GET" {
		t, erro:= template.ParseFiles("login.html")//文件要放在项目的目录中
		if erro==nil{
			t.Execute(w, nil)
		}else {
			log.Fatal(erro.Error())
		}
	} else {
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
		fmt.Fprint(w,"name:",r.Form["username"],"\npassword",r.Form["password"])
	}

}

func main() {
	http.HandleFunc("/login",login)
	err:=http.ListenAndServe(":9090",nil)
	if err!=nil {
		log.Fatal(err.Error())
	}
}

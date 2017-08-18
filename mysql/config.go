package main

import (
	"database/sql"
	"log"
	"fmt"
	_"github.com/go-sql-driver/mysql"
	"encoding/json"
	"net/http"
	"strconv"
)

func ReturnData(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if r.Method == "GET" {
		db:=connectSql()
		defer db.Close()
		fmt.Fprint(w, query1(db))
	}else {
		result:=UserData{State:300,Message:"请求方式错误   应该是get",Data:nil}
		b,_:=json.Marshal(result)
		fmt.Fprintf(w,string(b))
	}
}
func deleteData(w http.ResponseWriter,r *http.Request)  {
	r.ParseForm()
	if r.Method=="POST" {
		db:=connectSql()
		defer db.Close()
		reId:=r.FormValue("id")
		id,erro:=strconv.Atoi(reId)
		if erro!=nil {
			result:=UserData{State:300,Message:"id  不是数字",Data:nil}
			b,_:=json.Marshal(result)
			fmt.Fprintf(w,string(b))
			return
		}
		if 1==delect(db,id) {
			result:=UserData{State:200,Message:"删除成功",Data:nil}
			b,_:=json.Marshal(result)
			fmt.Fprintf(w,string(b))
			return
		}
	}else {
		fmt.Fprintf(w,"请求错误")
	}

}

func check(err error) {
	if err != nil {
		log.Fatal(err.Error())
		panic(err)
	}
}

func connectSql() *sql.DB {
	db, erro := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/go?charset=utf8")
	check(erro)
	//insert(db)
	//defer db.Close()
	return db
}
func main() {
	http.HandleFunc("/getuser",ReturnData)
	http.HandleFunc("/delete",deleteData)
	err:=http.ListenAndServe(":9090",nil)
	check(err)
}

type User struct {
	Id       int `json:"id"`
	Age      int `json:"age"`
	Addr     string `json:"addr"`
	Username string `json:"username"`
}

type Result struct {
	Data []User `json:"data"`
}

type UserData struct {
	State   int `json:"state"`
	Data    []map[string]string `json:"data"`
	Message string `json:"message"`
}

func query1(db *sql.DB) string {

	rows, erro := db.Query("SELECT * FROM user")
	check(erro)
	colum, erro := rows.Columns()
	check(erro)

	vals := make([][]byte, len(colum))

	scans := make([]interface{}, len(colum))
	//这是使用scan引用vals,把数据填充到  []byte中
	for k, _ := range vals {
		scans[k] = &vals[k]
	}
	result := UserData{State: 200, Message: "获取成功"}
	for rows.Next() {
		rows.Scan(scans...)
		row := make(map[string]string)
		for k, v := range vals {
			key := colum[k]
			row[key] = string(v)
		}
		result.Data = append(result.Data, row)
	}
	b, erro := json.Marshal(result)
	fmt.Println(string(b))
	return string(b)
}

func query(db *sql.DB) {
	rows, erro := db.Query("select *from user")
	check(erro)
	var jsondata Result
	for rows.Next() {
		var username string
		var id int
		var age int
		var addr string
		rows.Scan(&username, &id, &age, &addr)
		jsondata.Data = append(jsondata.Data, User{Username: username, Age: age, Id: id, Addr: addr})
	}
	fmt.Println(jsondata)
	b, err := json.Marshal(&jsondata)
	check(err)
	fmt.Println(string(b))
}

func updata(db *sql.DB) {
	stmt, erro := db.Prepare("UPDATE user SET addr=? WHERE id=?")
	check(erro)
	res, erro := stmt.Exec("四川从大股东", 2)
	check(erro)
	affect, erro := res.RowsAffected()
	check(erro)
	fmt.Println(affect)
}
func delect(db *sql.DB,id int) int64{
	stmt, erro := db.Prepare("DELETE FROM user where id=?")
	check(erro)
	res, erro := stmt.Exec(id)
	check(erro)
	affect, erro := res.RowsAffected()
	check(erro)
	fmt.Println(affect)
	return affect
}

func insert(db *sql.DB) {
	stmt, erro := db.Prepare("INSERT user SET username=?,age=?,addr=?")
	check(erro)
	res, erro := stmt.Exec("golang", "56", "司参股或记大过")
	check(erro)
	id, erro := res.LastInsertId()
	check(erro)
	fmt.Println(id)
}

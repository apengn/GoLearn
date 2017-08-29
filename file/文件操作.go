package main

import (
	"os"
	"fmt"
	"io/ioutil"
	"bufio"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

//判断文件是否存在  存在 true   不存在false
func checkFileIsExist(fileName string) bool {
	if _, erro := os.Stat(fileName); os.IsNotExist(erro) {
		return false
	}
	return true
}
func main() {
	filename := "file/cc1.text"
	var writeString = "测试数据"
	var f *os.File
	var erro error
	if checkFileIsExist(filename) {
		f, erro = os.OpenFile(filename, os.O_APPEND, 0666)
		fmt.Printf("文件存在")
	} else {
		f, erro = os.Create(filename)
		fmt.Printf("文件不存在，创建一个")
	}
	check(erro)
	defer f.Close()
	//写入文件方式  1
	n, err1 := f.WriteString(writeString)
	fmt.Printf("写入%d个字节n", n)
	//写入文件方式二 字节数组
	check(err1)
	var d1 = []byte("写入文件方式二")
	err1 = ioutil.WriteFile(filename, d1, 0666)
	check(err1)

}

func writeFile3(fileName string )  {
	f,erro:=os.Create(fileName)
	check(erro)
	defer f.Close()
	var d=[]byte("")
	n,erro:=f.Write(d)
	check(erro)
	fmt.Printf("写入%d字节数据",n)
	f.Sync()
}

func writeFile4(f *os.File){
	defer f.Close()
	w:=bufio.NewWriter(f)
	var d=[]byte("writeFile4")
	n,erro:=w.Write(d)
	check(erro)
	fmt.Printf("写入%d字节数据",n)
	w.Flush()

}

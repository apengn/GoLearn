package main

import (
	"fmt"
	"unsafe"
)

func main() {
	u:=new(user)
	fmt.Println(*u)
	pName:=(*string)(unsafe.Pointer(u))
	*pName="张三"
	pAge:=(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(u))+unsafe.Offsetof(u.age)))
	*pAge = 20

	fmt.Println(*u)
}
type user struct {
	name string
	age int
}
//func main() {
//	i:=10
//	ip:=&i
//
//	//var fp *float64=(*float64)(ip)
//	//unsafe.Pointer是一种特殊意义的指针，它可以包含任意类型的地址，有点类似于C语言里的void*指针，全能型的。
//	var fp  *float64=(*float64)(unsafe.Pointer(ip))
//	*fp=*fp*3
//	fmt.Println(i)
//}

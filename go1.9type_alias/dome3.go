package main

import "fmt"

func main() {

	a:=200

	b:=&a
	*b++
	c:=&b
	fmt.Println(*b) //取b的值
	fmt.Println(*c)//取c的值     =   b的值  而b的值是a的内存地址
	fmt.Println(&a)//取a 的内存地址
	fmt.Println(&b) //取b的内存地址
	fmt.Println(a)
}

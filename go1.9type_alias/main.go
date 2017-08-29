package main

import "fmt"

type User struct {

}
type myint1 int  //基于int类型创建了新类型myint1
type myint2 =User//创建了一个int的别名    myint2              注意：类型别名的定义是用“=”

//var i int = 1

//var i1 myint1=i  //抱错 因为 i是int类型的   i1是myint1类型的    不同类型的不能赋值
//var i1 myint1=myint1(i)  //可以进行 强制类型转换进行赋值

//var i2 myint2 = i
func (i myint1)m2()  {
	fmt.Println("myint1.m1")
}
func (i myint2)m2()  {
	fmt.Println("myint2.m2")
}

type I interface {
	m2()
}

type I1 I
type I2=I
func main() {
	var i1 I1=myint1(0)
	var i2 I2=myint1(0)
	var i I=myint1(0)

	i1=i2
	i1=i

	i1=i2
	i1=i
	i2=i
	i2=i1

//fmt.Println(i2,i1)
//	var i1  myint1
//	var i2  myint2
//    var iu User
//	var in I
//	var iuI I
//	in=i2
//	iuI=iu
//	i1.m1()
//	i2.m2()
//	in.m2()
//	iuI.m2()

}

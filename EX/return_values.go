package main

import "fmt"

func vals()(int,int)  {
	return 3,9
}

func valus()(a ,b int)  {
	a=1
	b=4
	return
}

func main() {
	a,b:=vals()

	fmt.Println(a,b)

	_,c:=vals()
	fmt.Println(c)
	sum(vals())
	sum([]int{99,77,7,6,4}...)
	sum(6,5,)
	sum(6,5,4,4,6,3,54)
}
//可变参数
func sum(nums ...int)  {
	fmt.Println(nums)
	total:=0
	for _,num:=range nums{
		total+=num
	}
	fmt.Println(total)
	fmt.Println(intSeq()())
	fmt.Println(intSeq()())
	fmt.Println(intSeq()())

	ints:=intSeq()

	fmt.Println(ints())
	fmt.Println(ints())
	fmt.Println(ints())
	fmt.Println(ints())
}
func intSeq()func() int  {
	i:=0
	return func() int {
		i+=1
		return i
	}
}




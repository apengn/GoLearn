package main

import "fmt"

func zeroval(ival int)  {//进行的是值的copy
	ival=0
}

func zeroptr(iptr *int)  {
	*iptr=0
}

func main() {
i:=1
	fmt.Println("initila:",i)
	zeroval(i)
	fmt.Println(i)
	zeroptr(&i)
	fmt.Println("zeroptr",i)
	fmt.Println("pointer:",&i)
}
package main

import "fmt"

func main()  {
	a:=7
	list:=[]int{a}
	lists(a,&list)
}

func lists(a int,list*[]int)  {
  fmt.Println(a)
}
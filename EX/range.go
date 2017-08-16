package main

import (
	"fmt"
)

func main() {
	nums:=[]int{2,3,4}

	sum:=0
	for _,num:=range nums{
		sum+=num
	}
	fmt.Println(sum)

	for i,num:=range nums{
		if num==3 {
			fmt.Println("index",i)
		}
	}

	kvs:=map[string]string{"a":"a1","a1":"a1","a2":"a2"}

	for k,v:=range kvs{
		fmt.Println(k,v)
	}
for k:=range kvs{
	fmt.Println(k)
}
	for i,c:=range "gogog"{
		fmt.Println(i,string(c))
	}

}

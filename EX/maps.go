package main

import "fmt"

func main() {



	m:=make(map[string]int)
	m["k1"]=8
	m["k2"]=3
	m["k3"]=86
	fmt.Println(m)


	v1:=m["k1"]
	fmt.Println(v1)
	fmt.Println(len(m))

	delete(m,"k1")
	fmt.Println(m)


	_,prs:=m["k12"]
	fmt.Println(prs)
	n:=map[string]int{"w1":2,"w2":221,"w":22}
	fmt.Println(n)
}

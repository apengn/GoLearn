package main

import "fmt"

func main() {

MyStuct{}.m1()
	MyStuct{}.m1()


}

func (t T1)m1()  {
	fmt.Println("T1.m1")
}
type T1 struct {

}
type T2=T1
type MyStuct struct {
	T2
	T1
}

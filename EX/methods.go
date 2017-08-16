package main

import "fmt"

type Rect struct {
	width,height int
}

func (r Rect)area()int  {
	return r.height*r.width
}
func (r *Rect)perim()int  {
	return 2*r.width+2*r.height
}

func main() {
	r:=Rect{10,5}

	fmt.Println("aren:",r.area())
	fmt.Println("perim:",r.perim())

	rp:=&r
	fmt.Println(rp.area())
	fmt.Println(rp.perim())
}

package main

import "fmt"

type Usb interface {
	Name() string
	Collected
}

type Phone struct {
	name string
}
type Collected interface {
	Collected()
}

type Collected1 interface {
	Collected1()
}
func (p Phone) Name()string  {
	return p.name
}
func (p Phone) Collected()  {
	fmt.Println("Collected:",p.name)
}
func (p Phone) Collected1()  {
	fmt.Println("Collected1:",p.name)
}
func main() {

	var usb  Collected1
	phone:=Phone{name:"phone"}
	usb=phone
	//usb.Collected1()
	is(usb)

}

func is(o interface{})  {

	//if c,ok:=o.(Usb);ok{
	//	c.Collected()
	//	fmt.Println(c.Name())
	//}else {
	//	c.Collected()
	//}

	switch v:=o.(type) {
	case Collected:
		v.Collected()
		fmt.Println("Collected")
	case Collected1:
		v.Collected1()
		fmt.Println("Collected1")
	case Usb:
		v.Collected()
		fmt.Println("Usb")
	default:
		fmt.Println("xxxxxxxx")
	}
}

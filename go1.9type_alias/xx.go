package main

import (
	"fmt"
	"reflect"
)

func increase(d int)(ret int)  {

	defer func() {
		ret++
		fmt.Println(ret)
	}()
	fmt.Println(d)

	fmt.Println("xxxxxxxxxx")
	return d
}
func main() {
  fmt.Println(reflect.TypeOf(""))
	fmt.Println(increase(1))
	//fmt.Println(fmt.Sprintf("%.2f",3/5.0))
  //fmt.Println(strconv.FormatFloat(0.0003/5.00,'E',2,64))
}

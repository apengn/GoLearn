package main

import (
	"fmt"
)

type Person struct {
	name string
	age int
}

func main() {



	fmt.Println(Person{"BOB",90})


	fmt.Println(Person{name:"Alice",age:65})

	fmt.Println(Person{name:"Fred"})

	fmt.Println(&Person{name:"Ann" ,age:40})

	s:=Person{name:"Sean",age:89}
	fmt.Println(s.name)

	sp:=&s
	fmt.Println(sp)
	sp.age=99
	fmt.Println(sp.age)
	fmt.Println(s.age)
	fmt.Println(*sp)




}
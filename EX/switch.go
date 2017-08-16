package main

import (
	"fmt"
	"time"
)

func main() {
	i:=2
	
	fmt.Println("Write",i,"as")

	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Saturday,time.Sunday:
		fmt.Println("Its the weekend")
	default:
		fmt.Println("Its a weekday")
	}


	t:=time.Now()

	switch  {
	case t.Hour()<12:
		fmt.Println("it's before no on")
	default:
		fmt.Println("it's  after n oon")
	}

	whatAmi:= func( i interface{}){
		switch t:=i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm a int")
		default:
			fmt.Printf("Don't know type %T\n",t)
		}
	}
	whatAmi(true)
	whatAmi(1)
	whatAmi("hey")


}

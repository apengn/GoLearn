package main

import (
	"fmt"
	
)

func main() {
	out:=make(chan int,1)
	out<-2
    fmt.Println(<-out)

}

package main

import (
	"time"
	"fmt"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		fmt.Println("--------------one")
		time.Sleep(time.Second * 4)
		fmt.Println("----------one----")
		c1 <- "one"
	}()
	go func() {
		fmt.Println("----------two----")
		time.Sleep(time.Second * 8)
		c2 <- "two"
		fmt.Println("---------two-----")
	}()

	for i := 0; i < 2; i++ {
		fmt.Println("=========")
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
		}
	}

}

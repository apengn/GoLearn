package main

import "fmt"

func main() {
	message:=make(chan string)

	go func() {
		message<-"ping"
	}()
	msg:=""
	msg =<-message
	fmt.Println(msg)

}

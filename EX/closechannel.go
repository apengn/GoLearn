package main

import (
	"fmt"
	"time"
)
func main() {
	jobs:=make(chan int,5)
	done:=make(chan bool)
	go func() {
		for  {
			j,more:=<-jobs
			if more{
				fmt.Println("received job",j)
			}else {
				fmt.Println("received all jobs")
				done<-true
				return
			}
		}
	}()
	for j:=0;j<=2 ;j++  {
		time.Sleep(time.Second)
		jobs<-j
		fmt.Println("sent job",j)
	}
	close(jobs)
	fmt.Println("sent all jobs")
	fmt.Println(<-done)
}
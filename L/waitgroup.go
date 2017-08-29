package main

import (
	"sync"
	"net/http"
	"fmt"
)

func main() {


	var wg sync.WaitGroup

	var urls =[]string{
		"http://www.golang.org/",
		"http://www.google.com/",
		"http://www.somestupidname.com/",
	}

	for _,url:=range urls{
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			res,_:=http.Get(url)
			fmt.Println(res.Status)
		}(url)
	}
	wg.Wait()
	fmt.Println("exit")
}

package main

import "fmt"

func main() {
	fmt.Println(fact(7))
}
func fact( n int)  int{
	if n==0{
		return 1
	}

	//7*6*5*4*3*2*1*1
	return fact(n-1)*n
}


//  face(7)=7*face(6)  5040
 //face(6)=6*face(5) 720
//face(5)=5*face(4) 120
//face(4)=4*face(3) 24
//face(3)=3*face(2) 6
//face(2)=2*face(1) 2
//face(1)=1*face(0) 1
//face(0)=1

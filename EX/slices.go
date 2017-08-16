package main

import "fmt"

func main() {

	s := make([]int, 3)

	fmt.Println(s)
	s = append(s, 4, 4, 35, 543, 4)
	s[1] = 99999
	fmt.Println(s)

	c := make([]int, len(s))

	copy(c, s)
	fmt.Println("c", c)

	l := c[2:7]
	fmt.Println("c[2:7]", l)

	l = s[:5]
	fmt.Println("sl2", l)

	t := []int{2, 5, 5, 3, 22}
	fmt.Println(t)
    t=append(t,9999)
	fmt.Println(t)

	twD:=make([][]int,3)

	for i:=0;i<3;i++ {
		innerLen:=i+1
		twD[i]=make([]int,innerLen)
		for j:=0;j<3;j++ {
			twD[i][j]=i+j
		}
	}

	fmt.Println(twD)
}

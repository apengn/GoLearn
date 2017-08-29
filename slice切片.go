package main

import "fmt"

func main() {

	var ss = make([]string, 0)//length:0 ,	addr:0x525b88,	isnil:false
	fmt.Printf("length:%v ,\taddr:%p,\tisnil:%v\n", len(ss), ss, ss == nil)
	for i:=0;i<10 ;i++  {
		ss=append(ss, fmt.Sprintf("s%d",i))
	}
	fmt.Println(ss)
	//目标切片大小小于源切片大小，copy就按照目标切片大小复制，不会报错。
	//var da=make([]string,0,10)//目标切片大小0，容量10，copy不能复制
	//var da=make([]string,10,10)//目标切片大小10，容量10，copy能复制
	//var da=make([]string,10,0)//目标切片大小10，容量0，copy  报错
	//var da=make([]string,10)//目标切片大小10，copy  copy能复制10ge
	var da=make([]string,5)//目标切片大小5，copy  copy能复制5ge
	cc:=copy(da,ss)
	fmt.Printf("copy to da (len=%d) \t copied：%d\t da :%v",len(da),cc,da)

}

func qiepian() {
	var ss []string //默认值 length:0 ,	addr:0x0,	isnil:true
	fmt.Printf("length:%v ,\taddr:%p,\tisnil:%v", len(ss), ss, ss == nil)

	for i := 0; i < 10; i++ {
		ss = append(ss, fmt.Sprintf("s%d", i))
	}
	fmt.Println()
	fmt.Printf("length:%v ,\taddr:%p,\tisnil:%v\n", len(ss), ss, ss == nil)
	fmt.Println(ss)
	fmt.Println(ss[0:2])

	index := 5
	//删除切片元素   增加以前的某个值
	ss = append(ss[:index], ss[index+1], ss[index+3])
	fmt.Println(ss)
	rear := append([]string{}, ss[index:]...)
	fmt.Println(rear)
	ss = append(ss, rear...)
	fmt.Println(ss)

}

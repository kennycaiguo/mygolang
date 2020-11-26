package main

import "fmt"

func main()  {
	//定义切片的方式1，跟数组类似，只不过不能写长度
	var str =[]string{"pussy","vagina","clitoris","labias"}
	for index,v:=range str{
		fmt.Printf("index:%d ,value:%s\n",index,v)
	}
	fmt.Println("=================================================")
	//定义切片的方式2，从数组中获取切片，因为切片是基于数组的
	var i [6]int =[6]int{1,11,22,33,2,5}
	s:=i[1:4] //从第一个开始到第四给结束，不包含第四个，也就是只取3个
	for index,v:=range s{
		fmt.Printf("index:%d ,value:%d\n",index,v)
	}
    //定义切片的方式3,特殊限定
    s2:=i[1:] //11,22,33,2,5
	for index,v:=range s2{
		fmt.Printf("index:%d ,value:%d\n",index,v)
	}
	fmt.Println("@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@")
	s3:=i[:4] //1,11,22,33，注意这里是不包含下标的那个
	for index,v:=range s3{
		fmt.Printf("index:%d ,value:%d\n",index,v)
	}
	fmt.Println("############################################")
	s4:=i[:] //1,11,22,33,2,5 这个相当于将整个数组赋值给切片
	for index,v:=range s4{
		fmt.Printf("index:%d ,value:%d\n",index,v)
	}
    fmt.Printf("length of i:%d ,capacity of i:%d\n", len(i), cap(i))
    fmt.Printf("length of s:%d ,capacity of s:%d\n", len(s), cap(s))
	s5:=s3[:3]
	fmt.Println(s5)
}

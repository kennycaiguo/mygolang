package main

import "fmt"

func main()  {
	a:=[]string{"pussy","vagina","breast","labias"}
	var b []string
	c:=make([]string,4,4)
	i := copy(b, a)
	fmt.Println(b,i) //[] 0
	i = copy(c,a) //使用copy函数，目的切片必须要有长度和容量否则不会拷贝内容，且和源切片一样，才能拷贝全部
	fmt.Println(c,i) //[pussy vagina breast labias] 4

}

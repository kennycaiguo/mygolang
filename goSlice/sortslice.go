package main

import (
	"fmt"
	"sort"
)

func main()  {
	s:=[...]int{5,4,3,2,1}
	s1:=s[:]
	sort.Ints(s1) //这个方法对切片进行升序排列（从小到大），并且没有返回值
	fmt.Println(s1)
}

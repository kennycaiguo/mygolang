package main

import (
	"fmt"
	"strings"
)

func main() {
	str:="how do you do ? i do will , how are you ? i am fine"
	//1.分字符串
	words :=strings.Split(str," ")
	//2.遍历切片存储到map
	m := make(map[string]int, 1)
	for _,w :=range words{
		//如果原来的map中不存在这个key，则m[w]=1
		if m[w]==0{
			m[w]=1
		}else {
			m[w]++
		}


	}
    //3.
    for key,value:=range m{
		fmt.Printf("%s 出现了%d次\n",key,value)
	}


}

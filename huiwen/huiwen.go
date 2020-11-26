package main

import "fmt"

func main() {
	//回文："上海自来水来自海上"
	//     “山西运煤车煤运西山”
	str:="东莞送货车货送莞东"
	//把字符串转为rune
	rstr:=[]rune(str)
	for i:=0;i<len(rstr)/2;i++{
		if rstr[i]!=rstr[len(rstr)-1-i]{
			 fmt.Println("不是回文")
			return
		}

	}
	fmt.Println("是回文")
}

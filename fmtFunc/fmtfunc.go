package main

import "fmt"

func main() {
	DoInfo()
}

func DoInfo() {
	var (
		name    string
		age     int
		married bool
	)
	fmt.Printf("请依次输入姓名，年龄，和婚否（true/false）")
	fmt.Scan(&name, &age, &married)
	fmt.Printf("扫描结果 name:%s age:%d married:%t \n", name, age, married)
}


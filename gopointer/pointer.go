package main

import "fmt"

func main() {
	num:=20
	fmt.Println(&num) //获取num的内存地址:0xc00000a0d0
    addr:=&num
    fmt.Println(*addr)//根据地址获取值：结果：20
}

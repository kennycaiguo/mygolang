package main

import "fmt"

//递归调用，即函数自己调用自己
func fibonari(x int) int{
	if x==0 {
		return 0;
	}
	if(x==1){
		return 1;
	}
	return fibonari(x-1)+fibonari(x-2)
}
//利用递归实现阶乘
func factorial(x int)(res int){

	if x>1{
		res = x*factorial(x-1)
		return res
	}
	return 1
}
func main() {
   /*for i:=0;i<10;i++{
   	  fmt.Printf("%d ",fibonari(i))//利用递归实现斐波拉数列，从第三给数字开始，每个数字是值是前两个数字的和
   }*/
	x:=5
	fmt.Printf("%d 的阶乘是%d\n",x,factorial(x))
}

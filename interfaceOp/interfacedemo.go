package main

import "fmt"

type Actioner interface {
	Say()
	Eat()

}

type cat struct {}
type dog struct {}

//实现接口的方法
func (c cat)Say()  {
	fmt.Println("miaomiaomiao...")
}

func (d dog)Say()  {
	fmt.Println("wangwangwang...")
}
func (c cat)Eat()  {
	fmt.Println("i want fish...")
}

func (d dog)Eat()  {
	fmt.Println("i want a bone...")
}

func main() {
     var a Actioner //声明一个接口变量
     c:=cat{} //声明一个cat结构体变量
     d:=dog{} //声明一个dog结构体变量
     a=c //将cat的对象赋值给x
     a.Say() //调用的是cat的Say()方法
     a=d //将dog的对象赋值给x
     a.Eat()//调用的是dog的Eat()方法


}

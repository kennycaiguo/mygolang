package main

import "fmt"

//接口与接口间可以通过嵌套创造出新的接口。
type speaker interface {
	Speak()
}
type mover interface {
	Move()
}

type animal interface {//结构体的嵌套，只有实现了所有接口的所有方法才是该结构体的类型
	speaker
	mover
}

type bird struct {
	name string
}
type people struct {
	name string
	gender string
}

//鸟类实现接口
func (b bird)Speak()  {
	fmt.Printf("%s 叽叽叽的叫\n",b.name)
}
func (b bird)Move()  {
	fmt.Printf("%s 在动\n",b.name)
}
//人类实现接口
func (p people)Speak()  {
	fmt.Printf("%s,%s 再说话。。。\n",p.name,p.gender)
}
func (p people)Move()  {
	fmt.Printf("%s,%s 在动\n",p.name,p.gender)
}

func main() {
	var a animal
	a = people{ //animal接口对象接收人类结构体的实例
		"jack",
		"male",
	}

	a.Speak()//人类的Speak方法
	a.Move() //人类的Move方法
	a = bird{ //animal接口对象接收鸟类结构体的实例
		"parrot",
	}
	a.Speak()//鸟类的Speak方法
	a.Move()//鸟类的Move方法
}
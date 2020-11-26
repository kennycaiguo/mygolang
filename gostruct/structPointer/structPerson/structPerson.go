package structPerson

import "fmt"

type Person struct {
	Name string
	Age int
	Gender string
	City string
}
//go语言中的“构造函数”
func NewPerson(name string,age int,gender,city string) *Person{ //go语言中的结构体是没有构造函数的，我们来模拟一个
	return &Person{
		name,age,gender,city,
	}
}

//在go语言中，方法就是指扩展方法，用来扩展对象的功能
func (p Person)Dream(){ //这就是一个方法，相当于c#中的扩展方法，可以用p.Dream()来调用
	fmt.Printf("%s做梦都想着学习go语言",p.Name);
}
package main

import (
	"fmt"
	"./structPerson"
)



func main() {
	//persontest1()
	//persontest2()
	//persontest3()
	//persontest4()
	//persontest5()
	persontest6()

}

func persontest6() {
	p6:=&person.Person{
		"Becken",
		20,
		"male",
		"zmt",
	}
	changesextofemale(*p6)//尝试将p6指向的结构体对象的性别修改
	fmt.Println(p6.Gender)//开始最终p6指向的结构体对象的性别并没有修改，可见go语言中结构体是值类型的
}

func changesextofemale(p person.Person) {
	p.Gender="female"
}

func persontest5() {
	p5:=&person.Person{
		"Becky",
		20,
		 "female",
		 "zmt",
	}
	fmt.Printf("p5 is :%#v\n",p5);
}

func persontest4() {
	p4:=&person.Person{
		Name:"Jackline",
		Age:20,
		Gender: "female",
		City: "dongguan",
	}
	//fmt.Printf("p4 is :%#v\n",p4) //p4 is :&main.structPerson{name:"Jackline", age:20, gender:"female", city:"dongguan"}
	fmt.Printf("p4 points to  :%#v\n",*p4)//p4 points to  :main.structPerson{name:"Jackline", age:20, gender:"female", city:"dongguan"}

}

func persontest3() { //使用键值对对结构体进行初始化
	p3 := person.Person{
		Name: "小王子",
		Gender: "male",
		City: "北京",
		Age:  18,
	}
	fmt.Printf("p5=%#v\n", p3) //p3=main.structPerson{name:"小王子", city:"北京", gender: "male",age:18}
}



func persontest1() {
	var p = new(person.Person)
	p.Name = "Jack"
	p.Gender = "male"
	p.Age = 20

	fmt.Printf("name:%s,age:%d,gender:%s", p.Name, p.Age, p.Gender)
}

func persontest2() {
	p2 := &person.Person{} //使用&对结构体进行取地址操作相当于对该结构体类型进行了一次new实例化操作。
	fmt.Printf("%T\n", p2)     //*main.structPerson
	fmt.Printf("p3=%#v\n", p2) //p3=&main.structPerson{name:"", age:0,gender=""}
	p2.Name = "xiaoming"
	p2.Age = 30
	p2.Gender="female"
	p2.City="beijing"
	//fmt.Println(p2)//&{xiaoming 30 female beijing}
	fmt.Println(*p2) //{xiaoming 30 female beijing}

}
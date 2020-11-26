package main

import "fmt"

//一个类型可以同时实现多个接口，而接口间彼此独立，不知道对方的实现。 例如，狗可以叫，也可以动。我们就分别定义Sayer接口和Mover接口，如下： Mover接口。

// Sayer 接口
type Sayer interface {
	say()
}

// Mover 接口
type Mover interface {
	move()
}
//dog既可以实现Sayer接口，也可以实现Mover接口。

type Dog struct {
	name string
}

// 实现Sayer接口
func (d Dog) say() {
	fmt.Printf("%s会叫汪汪汪\n", d.name)
}

// 实现Mover接口
func (d Dog) move() {
	fmt.Printf("%s会动\n", d.name)
}

func main() {
	var x Sayer
	var y Mover

	var a = Dog{name: "旺财"}
	x = a
	y = a
	x.say()
	y.move()
}

package main

import "fmt"

func main() {
	type test struct { //结构体占用一块连续的内存。
		a int8
		b int8
		c int8
		d int8
	}
	n := test{
		1, 2, 3, 4,
	}
	fmt.Printf("n.a %p\n", &n.a)//n.a 0xc00000a0c0
	fmt.Printf("n.b %p\n", &n.b)//n.b 0xc00000a0c1
	fmt.Printf("n.c %p\n", &n.c)//n.c 0xc00000a0c2
	fmt.Printf("n.d %p\n", &n.d)//n.d 0xc00000a0c3
}

package main

import "fmt"

func main() {
	var user struct{name string; age int}
	user.name="jack"
	user.age=18
	fmt.Println(user)
}

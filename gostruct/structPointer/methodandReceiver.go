package main

import (
	"fmt"
	"./structPerson"
)

func main() {
   p:=structPerson.NewPerson("Jack",18,"male","cincinati")
   fmt.Println(p)
   p.Dream()
}

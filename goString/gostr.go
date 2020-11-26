package main

import (
	"fmt"
	"unicode"
)
func main()  {
	count := 0
	s := "hello沙河小王子"
	for _, r := range s{
		if unicode.Is(unicode.Han, r) {
			count++
		}
	}
	fmt.Println("汉字个数：",count)
}

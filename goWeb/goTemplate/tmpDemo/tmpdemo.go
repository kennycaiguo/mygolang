package main

import (
	"fmt"
	"html/template"
	"os"
)

func main()  {
	tmpl,_:=template.ParseFiles("header.tmpl","content.tmpl","footer.tmpl")
	//tmpl.ExecuteTemplate(os.Stdout,"header",nil)
	//fmt.Println()
	tmpl.ExecuteTemplate(os.Stdout,"content",nil)
	//fmt.Println()
	//tmpl.ExecuteTemplate(os.Stdout,"footer",nil)
	fmt.Println()
	tmpl.Execute(os.Stdout,nil)
}

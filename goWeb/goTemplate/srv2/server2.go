package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"html/template"
	"io"
	"net/http"
)

//注意：body中的所有字段的第一个字母必须大写否则没有任何效果
type body struct {
	Title  string
	Content_header string
	Content_main string
}
type admin struct {
    Id int
    Username string
    Pwd string
}
func tmp(w http.ResponseWriter, r *http.Request)  {
	b:=body{
		"go template test",
		"这是内容标题",
		"欢迎来到首页",
		  //`<a href="/reg">注册</a><a href="/login">登陆</a>`,
	}
	tpl,err:=template.ParseFiles("index.html")


	if err!=nil{
		fmt.Println(err.Error())
	}

	tpl.Execute(w,b)

}

func userInfo(w http.ResponseWriter, r *http.Request){
	tpl,err:= template.ParseFiles("userinfo.html")
	if err!=nil{
		fmt.Println(err.Error())
	}
	db := initdb()
	// sqlstr:="select id,username,pwd from Admin where id=?"
	sqlstr := "select id,username,pwd from Admin where id>? "

	stmt, err := db.Prepare(sqlstr)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer stmt.Close()
	ad := new(admin)
	rows, err := stmt.Query(3)
	if err != nil {
		fmt.Println(err.Error())
	}
	for rows.Next() {
		err = rows.Scan(&ad.Id, &ad.Username, &ad.Pwd)
		tpl.Execute(w, ad)
	}
}

func initdb() *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/company?charset=utf8")
	if err != nil {
		fmt.Println(err.Error())
	}
	return db
}
func main()  {
	server :=http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/",tmp)
	http.HandleFunc("/user",userInfo)
	http.HandleFunc("/reg",regform)
	http.HandleFunc("/reghandler",reghandler)

	server.ListenAndServe()
}

func reghandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	a:=new(admin)
	a.Username = r.Form.Get("adname")
	a.Pwd =r.Form.Get("adpwd")
	//output:=fmt.Sprintf("admin name:%s ,pwd:%s",a.Username,a.Pwd)
	//io.WriteString(w,output)
	var db *sql.DB
   db=initdb()
   insertsql:="insert into Admin(username,pwd) values(?,?)"
   stmt,err:=db.Prepare(insertsql)
   if err!=nil{
   	fmt.Println(err.Error())
   }
   defer stmt.Close()
   _,err=stmt.Exec(a.Username,a.Pwd)
	if err!=nil{
		fmt.Println(err.Error())
	}
	io.WriteString(w,`注册成功，点击<a href="/">主页</a>进入主页`)
}

func regform(w http.ResponseWriter, request *http.Request) {
	tpl,err:= template.ParseFiles("reg.html")
	if err!=nil{
		fmt.Println(err.Error())
	}
	 tpl.Execute(w,nil)
}

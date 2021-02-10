package main

import (
	"fmt"
	_ "github.com/lib/pq"
	"github.com/jmoiron/sqlx"
	_ "github.com/go-sql-driver/mysql"

)

var db *sqlx.DB
var err error
type Client struct { //注意结构体的字段必须大写，否则sqlx会查询失败
	Id int
	Name string
	Balance int
}
func InitDb() (err error) {
	dsn:="root:root@tcp(127.0.0.1:3306)/gouser?charset=utf8mb4&parseTime=True"
	db,err= sqlx.Connect("mysql",dsn)
	if err != nil {
		fmt.Printf("connect DB failed, err:%v\n", err)
		return err
	}
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(10)
	return nil
}
func sqlxQueryone()  { //查询
	var c Client
	sql1:="select id,name,balance from account where id=?"
	err=db.Get(&c,sql1,4)
	if err!=nil{
		fmt.Printf("query failed ,err:%v",err)
		return
	}
	fmt.Printf("client:%v",c)
}
func queryxAll(){
	rows,err:=db.Queryx("select * from account")
	if err!=nil{
		fmt.Printf("queryx error:%v",err.Error())
		return
	}
	var c Client
	for rows.Next(){
		//rows.Scan(&c.Id,&c.Name,&c.Balance)
		rows.StructScan(&c)
		fmt.Println(c)
	}
}
func querytomap(){
	rows,err:=db.Queryx("select * from account")
	if err!=nil{
		fmt.Printf("queryx error:%v",err.Error())
		return
	}
	m:=make(map[string]interface{})
	for rows.Next(){
		rows.MapScan(m)
		fmt.Println(m)
	}
}
func querytoSlice()  {//一次性把所有的记录查询到一个Client切片中

	var clients []Client
	err := db.Select(&clients, "select * from account")
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	fmt.Printf("users:%#v\n", clients)
}
func insert(){
	sql:="insert into account(name,balance) values(?,?)"
	result, err:= db.Exec(sql, "Jackline", 2000)
	if err!=nil{
		fmt.Println(err.Error())
		return
	}
	id, err := result.LastInsertId()
	if err!=nil{
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("success,id inserted:%d",id)
}
func update()  {
	sql:="update account set balance=4000 where id=?"
	result, err:= db.Exec(sql, 1)
	if err!=nil{
		fmt.Println(err.Error())
		return
	}
	rowsAffected, err := result.RowsAffected()
	if err!=nil{
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("update successful,rows affected:%d",rowsAffected)
}
func del()  {
	sql:="delete from account where id=?"
	result, err := db.Exec(sql, 4)
	if err!=nil{
		fmt.Println(err.Error())
		return
	}
	rowsAffected, err := result.RowsAffected()
	if err!=nil{
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("delete successful,rows affected:%d",rowsAffected)
}
func main() {
   err:=InitDb()
   if err!=nil{
   	fmt.Printf("Init database failed",err)
   }
   //sqlxQueryone()
  // queryxAll()
   //querytomap()
   //querytoSlice()
   //insert()
   //update()
   del()
}

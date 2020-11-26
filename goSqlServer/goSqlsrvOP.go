package main

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/denisenkom/go-mssqldb"
)

type user struct {
	id int
	username string
	pwd string
}
func main(){
  prepareQuerySqlsrv()
}

//预编译查询，可以有效防止sql注入
func prepareQuerySqlsrv()  {

	//连接字符串
	connString := "server=.;database=kennyDb;user id=kenny;password=kenny1975;encrypt=disable"
	fmt.Println(connString)

	//建立连接
	conn, err := sql.Open("mssql", connString)
	if err != nil {
		log.Fatal("Open Connection failed:", err.Error())
	}
	defer conn.Close()

    //fmt.Println(conn)
	//产生查询语句的Statement
	stmt, err := conn.Prepare(`select * from [myuser]`)
	if err != nil {
		log.Fatal("Prepare failed:", err.Error())
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	defer rows.Close()
    //显示数据
	for rows.Next(){
		var u user
        err:=rows.Scan(&u.id,&u.username,&u.pwd)
        if err!=nil{
        	fmt.Println(err.Error())
			return
		}
       fmt.Printf("id:%d ,username:%s ,password:%s\n",u.id,u.username,u.pwd)
	}
}
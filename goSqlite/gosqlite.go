package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

const (
	dbDriverName = "sqlite3"
	dbName       = "e:/sqlite3/user.db3"
)

type user struct {
	id    int
	name  string
	pwd   string
	email string
}
func main(){
	//qrstr:="select * from user"
	db, err := sql.Open(dbDriverName, dbName)
	if err!=nil{
		fmt.Println(err.Error())
	}
	// prepareQuerySqlite(db,qrstr)
	/*var u user
	u.id=4
	u.name="Jolly"
	u.pwd="123"
	u.email="Jolly321@gmail.com"
	prepareInsert(db,u)*/

	/*var u user
	u.id=4
	u.name="Bobby"
	u.pwd="456"
	u.email="Jolly321@gmail.com"
	prepareUpdate(db,u)*/
	prepareDel(db,4)
}

func prepareQuerySqlite(db *sql.DB,query string) {

	stmt, err := db.Prepare(query)
	if err != nil {
		fmt.Printf("prepare failed, err:%v\n", err)
		return
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	defer rows.Close()
	for rows.Next(){
      var u user
      err:=rows.Scan(&u.id,&u.name,&u.pwd,&u.email)
      if err!=nil{
      	fmt.Println(err.Error())
		  return
	  }
      fmt.Printf("id:%d, username :%s, password:%s ,email: %s\n", u.id,u.name,u.pwd,u.email)
	}
}

func prepareInsert(db *sql.DB,u user)  {
	insertsql:="insert into user(id,name,pwd,email) values(?,?,?,?)"
	stmt,err :=db.Prepare(insertsql)

	if err!=nil{
		fmt.Println(err.Error())
	}
	defer stmt.Close()
	_,err=stmt.Exec(u.id,u.name,u.pwd,u.email)
	if err!=nil{
		fmt.Println(err.Error())
	}
	fmt.Println("插入数据成功")
}

func prepareUpdate(db *sql.DB,u user){
	updatesql:="update user set name=?,pwd=?,email=? where id=?"
	stmt,err:=db.Prepare(updatesql)
	if err != nil{
		fmt.Println(err.Error())
	}
	defer stmt.Close()
	_,err = stmt.Exec(u.name,u.pwd,u.email,u.id)
	if err!=nil{
		fmt.Println(err.Error())
	}
	fmt.Println("更新数据成功")
}

func prepareDel(db *sql.DB,id int){
	delsql:="delete from user where id=?"

	stmt,err :=db.Prepare(delsql)
	if err!=nil{
		fmt.Println(err.Error())
	}
	defer stmt.Close()
	_,err = stmt.Exec(id)
	if err!=nil{
		fmt.Println(err.Error())
	}
	fmt.Println("删除数据成功")
}
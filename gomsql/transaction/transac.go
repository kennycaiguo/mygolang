package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Account struct {
	id int
	name string
	balance int
}
var db *sql.DB
func InitDb(){
	dsn:="root:root@tcp(127.0.0.1:3306)/gouser?charset=utf8mb4&parseTime=True"
	db, _= sql.Open("mysql", dsn)

	err:= db.Ping()
	if err!=nil{
		fmt.Println(err.Error())
	}
}

func preparedBatchInsert()  {
	var m =map[string]int{
		"Amy":3000,
		"Alex":4000,
		"Joicy":2000,
		"Fred":2500,
		"Daisy":2800,
	}
	sqlstr:=`insert into account(name,balance) values(?,?)`
	stmt, err := db.Prepare(sqlstr)
	if err!=nil{
		fmt.Println(err.Error())
	}
	for k,v:=range m{
		stmt.Exec(k,v)
	}
}
/**

事务处理：如从Alex的账户里面给Joicy转账1000元
a:4000-3000
j:2000-3000
sql1="update account set balance=3000 where name=?"
 */
func doTransAct(){
	tx, err := db.Begin()
	if err!=nil{
		if tx!=nil{
			tx.Rollback()
		}
	}
	sql1:="update account set balance=balance-1000 where name=?"
	res1, err := tx.Exec(sql1, "Alex")
	if err!=nil{
		tx.Rollback()
		fmt.Printf("exec sql1 failed, err:%v\n", err)
		return
	}
	rowsAff1, _:= res1.RowsAffected()
	sql2:="update account set balance=balance+1000 where name=?"
	res2, err := tx.Exec(sql2, "Joicy")
	if err!=nil{
		tx.Rollback()
		fmt.Printf("exec sql2 failed, err:%v\n", err)
		return
	}
	rowsAff2, _:= res2.RowsAffected()
	if rowsAff1==1 && rowsAff2==1{
		tx.Commit()
		fmt.Println("提交事务啦。。。。")
	} else {
		tx.Rollback()
		fmt.Println("提交回滚啦。。。。")
	}
	fmt.Println("exec trans success!")
}
func main() {
	fmt.Println("======================begin================")
	InitDb()
	//preparedBatchInsert()
	doTransAct()
}
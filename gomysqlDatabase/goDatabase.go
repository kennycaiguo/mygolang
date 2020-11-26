package main

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/go-sql-driver/mysql"
)

type person struct {
	id int
	name string
	age int
	email string
}
func main()  {
	//dataSourceName的格式：用户名:密码@tcp(127.0.0.1:端口号)/数据库名称?charset=utf8
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/contact?charset=utf8")
	if err != nil{
		fmt.Println(err)
	}
	//readContact()
	//insertData(db)
	//updateData(db,"charlie",17)
	//deleteData(db,16)
	//prepareQuery(db,18)
	preparedInsert(db)
}

func readContact(db *sql.DB) {

	rows,err:=db.Query("select * from  structPerson where age>?",18)
	if err!=nil{
		log.Fatal(err)
	}
	for rows.Next(){
		var p person
		err:=rows.Scan(&p.id,&p.name,&p.age,&p.email)
		if err!=nil{
			fmt.Println(err.Error())
			return
		}
		fmt.Printf("id:%d name:%s age:%d email:%s\n",p.id, p.name, p.age,p.email)
	}
}

func insertData(db *sql.DB){

	sqlstr:="insert into structPerson(name,age,email) values(?,?,?)"

	ret,err:=db.Exec(sqlstr,"fanny",30,"fn@pussy.net")
	if err!=nil{
		fmt.Println(err.Error())
		return
	}
	retId,err:= ret.LastInsertId()

	if err!=nil{
		fmt.Println(err.Error())
	}

	fmt.Printf("insert success, the id is %d.\n", retId)
}

func updateData(db *sql.DB,name string,id int){
	sqlstr:="update structPerson set name=? where id=?"

	ret,err:=db.Exec(sqlstr,name,id)
	if err!=nil{
		fmt.Println(err.Error())
		return
	}
	n,err:=ret.RowsAffected()
	if err!=nil{
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("更新成功,受影响的行数：%d",n)
}

func deleteData(db *sql.DB,id int)  {
	delstr:="delete from structPerson where id=?"

	ret,err:=db.Exec(delstr,id)
	if err!=nil{
		fmt.Println(err.Error())
		return
	}
	n,err:=ret.RowsAffected()
	if err!=nil{
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("删除数据成功，受影响的行数：%d",n)
}

// 预处理查询示例
/*
预处理执行过程：

把SQL语句分成两部分，命令部分与数据部分。
先把命令部分发送给MySQL服务端，MySQL服务端进行SQL预处理。
然后把数据部分发送给MySQL服务端，MySQL服务端对SQL语句进行占位符替换。
MySQL服务端执行完整的SQL语句并将结果返回给客户端。
为什么要预处理？
优化MySQL服务器重复执行SQL的方法，可以提升服务器性能，提前让服务器编译，一次编译多次执行，节省后续编译的成本。
避免SQL注入问题。

*/
func prepareQuery(db *sql.DB,id int) {
	sqlStr := "select id, name, age ,email from structPerson where id > ?"
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Printf("prepare failed, err:%v\n", err)
		return
	}
	defer stmt.Close()
	rows, err := stmt.Query(id)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	defer rows.Close()
	// 循环读取结果集中的数据
	for rows.Next() {
		var p person
		err := rows.Scan(&p.id, &p.name, &p.age,&p.email)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return
		}
		fmt.Printf("id:%d name:%s age:%d,email:%s\n", p.id, p.name, p.age,p.email)
	}
}

func preparedInsert(db *sql.DB){
	insertsql:="insert into structPerson(name,age,email) values(?,?,?)"

	stmt,err:=db.Prepare(insertsql)
	if err!=nil{
		fmt.Println(err.Error())
		return
	}
	defer stmt.Close()
	_,err=stmt.Exec("Jolly",20,"Jolly@gmail.com")
	if err!=nil{
		fmt.Println(err.Error())
		return
	}
	//一条预编译sql了、可以执行多次
	_,err=stmt.Exec("Molly",21,"Mollygirl@gmail.com")
	if err!=nil{
		fmt.Println(err.Error())
		return
	}
	fmt.Println("插入数据成功")
}
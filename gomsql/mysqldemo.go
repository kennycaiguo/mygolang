package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)
//定义一个全局变量保存数据库连接池对象
//var db *sql.DB
/** 建表ddl
 create table `user` (
        `id` bigint(20) not null  auto_increment,
        `name` varchar(50) default ' ',
         `age`  int(4) default '18',
         `email` varchar(100) default ' ',
          primary key(`id`)
)engine=InnoDB auto_increment=1 default charset=utf8mb4;
 */
//这个用于接收查询数据
type User struct {
	id    int
	name  string
	age   int
	email string
}

func initDb() (*sql.DB,error) {
	dsn := "root:root@tcp(127.0.0.1:3306)/gouser?charset=utf8mb4&parseTime=True"
	db, err := sql.Open("mysql", dsn) //不会校验用户名密码是否正确,这里只验证dsn的格式是否正确
	if err != nil {
		fmt.Println("数据库格式错误")
		return nil,err
	}
	err = db.Ping()
	if err != nil {
		return nil,err
	}
	return db,nil
}


func insert(db *sql.DB,user User){
	strsql:="insert into user(name,age,email) values(?,?,?)";
	ret, err := db.Exec(strsql, user.name, user.age, user.email)
    if err!=nil{
    	fmt.Printf("insert fail,error:%#v",err)
		return
	}
	lastInsertId, _ := ret.LastInsertId()
	fmt.Printf("插入数据成功，id为%v",lastInsertId)

}
func preparedInsert(db *sql.DB,user User){
	strsql:="insert into user(name,age,email) values(?,?,?)";
	stmt,_:=db.Prepare(strsql)
	ret, err := stmt.Exec(user.name, user.age, user.email)
    if err!=nil{
    	fmt.Printf("insert fail,error:%#v",err)
		return
	}
	lastInsertId, _ := ret.LastInsertId()
	fmt.Printf("插入数据成功，id为%v",lastInsertId)

}
func queryOne(db *sql.DB,id int)(User)  {
	var user User
	sql:="select id,name,age,email from user where id=?"
	err := db.QueryRow(sql, id).Scan(&user.id, &user.name, &user.age, &user.email)
	if err!=nil{
		fmt.Printf("查询失败，错误为：%#v",err)
	}
	return user
}
func preparedQueryone(db *sql.DB,id int){
	querystr:="select id,name,age,email from user where id>?"
	stmt,err:=db.Prepare(querystr)
	if err!=nil{
		fmt.Printf("prepare failed ,err:%v",err)
		return
	}
	rows, err := stmt.Query(1)
	if err!=nil{
		fmt.Printf("query failed ,err:%v",err)
		return
	}
	var u User
	defer rows.Close()
	for rows.Next(){
		 rows.Scan(&u.id, &u.name, &u.age, &u.email)
		 fmt.Printf("user :%v\n",u)
	}
}
func queryManyGreater(db *sql.DB,id int)(*sql.Rows){
	sql:="select id,name,age,email from user where id>?"
	rows, err := db.Query(sql, id)
	if err!=nil{
		fmt.Printf("query failed ,err:%#v",err)
	}
	return rows
}
func queryManyLesser(db *sql.DB,id int)(*sql.Rows){
	sql:="select id,name,age,email from user where id<?"
	rows, err := db.Query(sql, id)
	if err!=nil{
		fmt.Printf("query failed ,err:%#v",err)
	}
	return rows
}
func showRows(rows *sql.Rows){
	var u User
	for rows.Next(){
       rows.Scan(&u.id,&u.name,&u.age,&u.email)
       fmt.Printf("user:%#v\n",u)
	}
	rows.Close()
}

func updateById(db *sql.DB,user User,id int){
	sqlstr:="update user set name=?,age=?,email=? where id=?"
	_, err := db.Exec(sqlstr, user.name, user.age, user.email, user.id)
     if err!=nil{
     	fmt.Printf("update failed err:%#v",err)
		 return
	 }
	 fmt.Println("update successful")
}
func deleteById(db *sql.DB,id int)  {
	delsql:="delete from user where id=?"
	_, err := db.Exec(delsql, id)
	if err!=nil{
		fmt.Printf("delete failed ,err:%#v",err)
		return
	}
	fmt.Println("delete successfully!!!")
}
func main() {
	db,err := initDb()
	if err != nil {
		fmt.Printf("connect database failed,err:%v", err)
		return
	}
	//插入数据
   //user:=User{001,"jack",18,"Jack123@gmail.com"}
   //user:=User{004,"Ben",18,"Benabc@gmail.com"}
   //insert(db,user)
  //查询一行
 /* user:=queryOne(db,2)
  fmt.Printf("id=%d,name=%s,age=%d,email=%s",user.id,user.name,user.age,user.email)*/
   //showRows(queryManyGreater(db,0))
   //更新数据
	//user:=User{001,"jack",28,"HelloJack@gmail.com"}
	//updateById(db,user,1)
	//删除id为4的记录
	//deleteById(db,5)
	//预处理程序
	//preparedQueryone(db,1)
	//预处理插入
	user:=User{007,"Stacy",18,"Mikybbcc123@gmail.com"}
	preparedInsert(db,user)

}

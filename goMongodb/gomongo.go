package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
)

type  user struct {
	id  int
	username string
	password string
	email string
}

const URL string = "127.0.0.1:27017"

var c *mgo.Collection
var session *mgo.Session

func main()  {
	session, _ = mgo.Dial(URL)
	//切换到数据库
	db := session.DB("itcast")
	//切换到collection
	c = db.C("user")

    u:= new(user)
	u.id = 3
    u.username ="Jassica"
    u.password="123"
    u.email="Jassica@bigpussy.net"
    addnew(c,*u)
}

func addnew(c *mgo.Collection , u user){
	/*err:=c.Insert(u)
	if err == nil {
		fmt.Println("插入成功")
	} else {
		fmt.Println(err.Error())
		defer panic(err)
	}*/
	err:=c.Insert(map[string]interface{}{"id": u.id, "username": u.username, "password": u.password,"email":u.email})
	if err!=nil{
		fmt.Println(err.Error())
	}
	fmt.Println("插入数据成功！")
}



func update(c *mgo.Collection , u user)  {

}
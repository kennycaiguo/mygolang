package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type  usr struct {
	id  int
	username string
	password string
	email string
}

var url string = "mongodb://127.0.0.1:27017/"

func main()  {
	var (
		client *mongo.Client
		err error
		//result *mongo.InsertOneResult
	)
	// 建立mongodb连接
	clientOptions := options.Client().ApplyURI(url)
	if client, err = mongo.Connect(context.TODO(),clientOptions); err != nil {
		return
	}
	// 2, 选择数据库my_db
	db := client.Database("itcast")

	// 3, 选择表my_collection
	collection := db.Collection("user")
	u:=usr{
		4,
		"Miky",
		"123",
		"miky123@gmail.com",
	}


	insertOne(collection,u)
}

func insertOne(c *mongo.Collection,u usr){
	 result, err := c.InsertOne(context.TODO(), u);
	 if err != nil {
		fmt.Println(err)
		return
	}

	//fmt.Println(result)
	docId := result.InsertedID.(primitive.ObjectID)
	fmt.Println("自增ID:", docId.Hex())
}
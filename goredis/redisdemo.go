package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

var rdsclient *redis.Client

func initClient() (err error) {
	rdsclient=redis.NewClient(&redis.Options{
		Addr:"localhost:6379",
		Password: "",
		DB: 0,
		PoolSize: 100,
	})
	ctx, _ := context.WithTimeout(context.Background(), 40*time.Second)
	//defer cancel()​
	_, err= rdsclient.Ping(ctx).Result()
	//_, err= rdsclient.Ping(nil).Result() //报错
	 if err!=nil{
	 	fmt.Println(err.Error())
	 }
	 return err
}
func getkey(key string){
	ctx:=context.Background()
	val := rdsclient.Get(ctx, key)
	fmt.Printf("value:%v\n",val.Val())//注意:通过返回值的Val()方法才能得到需要的数据

}
func setKey(value interface{})  {
	ctx:=context.Background()
	cmd := rdsclient.Set(ctx, "age", value, 0)
	result, err := cmd.Result()
	if err!=nil{
		fmt.Printf("error:%v",err)
	}
	fmt.Printf("result:%v",result)
}
/**
127.0.0.1:6379> zrange rank 0 10 利用这个命令查看ZAdd是否成功
1) "Golang"
2) "Python"
3) "JavaScript"
4) "Java"
5) "C/C++"
 */
func zaddTest(){
	ctx:=context.Background()
	zsetKey := "rank"
	languages := []*redis.Z{
		&redis.Z{Score: 90.0, Member: "Golang"},
		&redis.Z{Score: 98.0, Member: "Java"},
		&redis.Z{Score: 95.0, Member: "Python"},
		&redis.Z{Score: 97.0, Member: "JavaScript"},
		&redis.Z{Score: 99.0, Member: "C/C++"},
	}
	// ZADD

	_, err := rdsclient.ZAdd(ctx,zsetKey, languages...).Result()
	if err != nil {
		fmt.Printf("zadd failed, err:%v\n", err)
		return
	}
	fmt.Println("ZAdd完成。。。")
}
func zincrbyTest()  {
	ctx:=context.Background()
	zsetKey := "rank"
	result, err := rdsclient.ZIncrBy(ctx, zsetKey, 5.0, "Golang").Result()//返回增加后的结果
	if err!=nil{
		fmt.Println(err.Error())
	}
	fmt.Printf("result：%v",result)
}

func getMaxScore(){
	ctx:=context.Background()
	zsetKey := "rank"
	// 取分数最高的3个
	ret, err := rdsclient.ZRevRangeWithScores(ctx,zsetKey, 0, 2).Result()//根据分数降序排列
	if err != nil {
		fmt.Printf("zrevrange failed, err:%v\n", err)
		return
	}
	for _, z := range ret {
		fmt.Println(z.Member, z.Score)
	}

}
//获取指定分数范围的记录
func getRangeScore()  {
	ctx:=context.Background()
	zsetKey := "rank"
	//取95到100的
	opt:=redis.ZRangeBy{
		Min:"97",
		Max:"100",
	}
	result, err := rdsclient.ZRangeByScoreWithScores(ctx, zsetKey, &opt).Result()
	if err!=nil{
		fmt.Printf(err.Error())
		return
	}
	fmt.Println("score list:\n")
	for _,z:=range result{
		fmt.Printf("%v:%v\n",z.Member,z.Score)
	}
}

func main() {
   err:=initClient()
   if err!=nil{
   	fmt.Println(err.Error())
   }
   fmt.Printf("连接redis成功\n")
   //getkey("name")
  // setKey(35)
  //getkey("age")

  //zaddTest()
  //zincrbyTest()
  //getMaxScore()
  getRangeScore()
}

package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	/*fmt.Println(now.Date())//今天日期 2020 November 5
	fmt.Println(now.Year()) //今天多少好 2020
	fmt.Println(now.Month()) //今天多少好 November
	fmt.Println(now.Day()) //今天多少好 5
	fmt.Println(now.Weekday()) //星期几 Thursday
	fmt.Println(now.Clock()) //下载几点 几分 几秒20 1 16
	fmt.Println(now.Hour()) //下载几点 20
	fmt.Println(now.Minute()) //下载几份 1
	fmt.Println(now.Second()) //下载几秒 16*/
	fmt.Println(now.Unix())//返回一个Unix时间戳：1604628196
     timestmp:=now.Unix()
	timeObj := time.Unix(timestmp, 0)//将时间戳转为时间：2020-11-05 20:05:32 -0600 CST
	fmt.Println(timeObj)
     later:=time.Now().Add(time.Hour)
     fmt.Println(later)//2020-11-05 21:08:06.4963053 -0600 CST m=+3600.031964901
     fmt.Println(later.Hour()) //21
	fmt.Println(now.Format("15:04 2006/01/02")) //格式化时间：20:28 2020/11/05
	fmt.Println(now.Add(time.Hour*(-24)).Format("2006-01-02 03:04:05.000 PM Mon Jan"))// 格式化昨天的时间：2020-11-04 08:29:44.204 PM Wed Nov

	//laterH:=time.Now().Add(time.Hour*2) //现在的时间加2个小时的写法

	//fmt.Println(laterH.Hour()) //21
	//fmt.Println(now.Format("yyyy-mm-DD HH:MM:SS"))//不是这样子写格式，而是参照go语言的出生时间格式
	//location, _ := time.LoadLocation("Asia/Shanghai")
	/*timestr:="2017/08/04 14:15:20"
	//timeObj, _ = time.ParseInLocation("2006/01/02 15:04:05",timestr , location)
	timeObj, _ = time.Parse("2006/01/02 15:04:05",timestr )
	fmt.Println(timeObj.Date())*/
	nextyear:="2021/01/01 14:15:20"
	//timeObj, _ = time.ParseInLocation("2006/01/02 15:04:05",timestr , location)
	nextObj, _ := time.Parse("2006/01/02 15:04:05",nextyear )
	fmt.Println(nextObj.Sub(now)) //time.Sub(u Time) 方法计算两个时间的间隔

}

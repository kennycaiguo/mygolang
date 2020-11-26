package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"reflect"
	"strconv"
	"strings"
)

type MysqlConf struct {
 Url string	`ini:"url"`
 Port int `ini:"port"`
 UserName string `ini:"username"`
 Password string `ini:"password"`
 Test bool `ini:"test"`
}

type RedisConf struct {
	Host string `ini:"host"`
	Port int `ini:"port"`
	Password string `ini:"password"`
	Database int `ini:"database""`
}

type Config struct {
	MysqlConf `ini:"mysql"`
	RedisConf `ini:"redis"`
}

func loadIni(file string,data interface{}) (err error) {
	//o.参数的校验
	t:=reflect.TypeOf(data)
	//0.1 data 必须是指针类型的，否则报错//
	if t.Kind()!=reflect.Ptr{
       err=errors.New("data参数必须是指针类型！！！")
       return  err
	}
	//0.2 data 必须是结构体指针类型的，否则报错
	if t.Elem().Kind()!=reflect.Struct{
		err=errors.New("data参数必须是结构体类型")
		return  err
	}
	//1.读取文件
	b, err:= ioutil.ReadFile(file) //一次读取文件的所有内容
		if(err!=nil){
			fmt.Printf("Open file failed error:%v",err.Error())
			return  err
		}
	var structname string
	lines := strings.Split(string(b), "\r\n") //利用换行符切割文件内容
	//2，需要逐行读取文件内容并且处理

	for idx,line:=range lines{
		//去掉字符串首尾的空格
		line := strings.TrimSpace(line)
		if len(line)==0{//遇到空行，直接跳过
			continue
		}

		//2.1如果是注释，直接跳过,以;或者#开头的是注释
		if strings.HasPrefix(line,";")||strings.HasPrefix(line,"#"){
			continue
		}
		//2.2 如果是[开头就是节（section）
		section:=line[1:len(line)-1]
		if strings.HasPrefix(line,"["){
			if line[0]!='[' || line[len(line)-1]!=']'{
				err = fmt.Errorf("line %d syntex error!!!",idx+1)
				return
			}
			if len(strings.TrimSpace(section))==0{
				err = fmt.Errorf("line %d syntex error!!!",idx+1)
				return
			}
			//根据section名称找到对应的结构体进行赋值
           // v:=reflect.ValueOf(data)
            for i:=0;i<t.Elem().NumField();i++{
              field:=t.Elem().Field(i)
             if section== field.Tag.Get("ini"){//找到了对应结构体
             	structname = field.Name

			 }
			}
		}else{
			//3.如果不是注释也不是节，就以=分割
			//3.1.进行分割，左边是key，右边是value
			if strings.Index(line,"=")==-1||strings.HasPrefix(line,"="){
				err = fmt.Errorf("line %d syntex error!!!",idx+1)
				return
			}
			index:=strings.Index(line,"=")
			key:=strings.TrimSpace(line[:index])
			value:=strings.TrimSpace(line[index+1:])
			//3.2.根据structname去找对应的嵌套结构体
			v:=reflect.ValueOf(data)
			svalue:=v.Elem().FieldByName(structname)//获取嵌套结构体的值信息
			stype:=svalue.Type() //获取嵌套结构体的类型信息
			if stype.Kind()!=reflect.Struct{
				err=fmt.Errorf("data中的%s字段需要是结构体类型的",structname)
				return err
			}
			var fieldName string
			var fieldType reflect.StructField
			//3.3遍历结构体的每一个字段，判断tag等不等于key
			for i:=0;i<svalue.NumField();i++{
				field:=stype.Field(i)//Tag信息是存储在类型信息中
				fieldType = field
               if field.Tag.Get("ini")==key{
                  //找到对应字段
               	fieldName = field.Name
                break
			   }

			}
			//3.4如果key==tag，给这个字段赋值

			//3.4.1 根据fieldName获取这个字段
			if len(fieldName)==0 {
				continue
			}
			obj:=svalue.FieldByName(fieldName)
			//3.4.2 对其赋值
			switch fieldType.Type.Kind() {
			case reflect.String:
              obj.SetString(value)
			case reflect.Int,reflect.Int8,reflect.Int16,reflect.Int32,reflect.Int64:
				var valueInt int64
				valueInt,err =strconv.ParseInt(value,10,64)
				if err!=nil{
					err=fmt.Errorf("line:%d data format error",idx+1)
					return
				}
				obj.SetInt(valueInt)
			case reflect.Bool:
				var boolValue bool
				boolValue, err = strconv.ParseBool(value)
				if err!=nil{
					err=fmt.Errorf("line:%d data format error",idx+1)
					return
				}
				obj.SetBool(boolValue)
			case reflect.Float32,reflect.Float64:
				var floatValue float64
				floatValue, err = strconv.ParseFloat(value,64)
				if err!=nil{
					err=fmt.Errorf("line:%d data format error",idx+1)
					return
				}
				obj.SetFloat(floatValue)
			}
			}


		}

    return err
}

func main() {
	var cfg = Config{}
   err:=loadIni("conf.ini",&cfg)
   if err!=nil{
   	fmt.Println(err.Error())
   }
   fmt.Printf("%#v\n",cfg)
}

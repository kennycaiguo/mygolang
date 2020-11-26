package main

import (
	"gopkg.in/gomail.v2"
	"strconv"
	"log"
	"fmt"
)

func SendMail(mailTo []string, subject string, body string) error {

	mailConn := map[string]string{
		"user": "caivenus2019@gmail.com", //不要使用我在主账号kenny887@gmail.com来实验，以免密码泄露
		"pass": "Venus2003",
		"host": "smtp.gmail.com",
		"port": "465",
	}

	port, _ := strconv.Atoi(mailConn["port"]) //转换端口类型为int

	m := gomail.NewMessage()

	m.SetHeader("From",  m.FormatAddress(mailConn["user"], "Venus")) //这种方式可以添加别名，即“XX官方”
	//说明：如果是用网易邮箱账号发送，以下方法别名可以是中文，如果是qq企业邮箱，以下方法用中文别名，会报错，需要用上面此方法转码
	//m.SetHeader("From", "FB Sample"+"<"+mailConn["user"]+">") //这种方式可以添加别名，即“FB Sample”， 也可以直接用<code>m.SetHeader("From",mailConn["user"])</code> 读者可以自行实验下效果
	//m.SetHeader("From", mailConn["user"])
	m.SetHeader("To", mailTo...)    //发送给多个用户
	m.SetHeader("Subject", subject) //设置邮件主题
	m.SetBody("text/html", body)    //设置邮件正文

	d := gomail.NewDialer(mailConn["host"], port, mailConn["user"], mailConn["pass"])

	err := d.DialAndSend(m)
	return err

}
func main() {
	//定义收件人
	mailTo := []string{
		"kennycai8888@outlook.com",
		"guimaihong@gmail.com",
	}
	//邮件主题为"Hello"
	subject := "Hello by golang gomailWithAttachment from smtp.gmail.com"
	// 邮件正文
	body := "Hello,by venus,how are u today?"

	err := SendMail(mailTo, subject, body)
	if err != nil {
		log.Println(err)
		fmt.Println("send fail")
		return
	}

	fmt.Println("send successfully")

}
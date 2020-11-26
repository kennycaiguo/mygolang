package main

import (
	"fmt"
	"gopkg.in/gomail.v2"
)

func main() {
	m := gomail.NewMessage()
	m.SetHeader("From", "caivenus2019@gmail.com")                     //发件人
	m.SetHeader("To", "kennycai8888@outlook.com")           //收件人
	m.SetAddressHeader("Cc", "guimaihong@gmail.com", "test")     //抄送人
	m.SetHeader("Subject", "Hello!")                     //邮件标题
	m.SetBody("text/html", "使用Go测试发送邮件!")     //邮件内容
	m.Attach("cmtp2.jpeg")       //邮件附件
	m.Attach("att.docx")       //邮件附件
	m.Attach("me.mp4")       //邮件附件

	d := gomail.NewDialer("smtp.gmail.com", 465, "caivenus2019@gmail.com", "Venus2003")
	//邮件发送服务器信息,使用授权码而非密码
	if err := d.DialAndSend(m); err != nil {
		fmt.Println("发送失败！！！")
	}else{
		fmt.Println("发送成功！！！")
	}
}

package core

import (
	"fmt"
	"gopkg.in/gomail.v2"
	"time"
)

func SendMail(content string,imageUrl string) (err error) {
	m := gomail.NewMessage()
	m.SetHeader("From", "rentmaterial@163.com")                     //发件人
	m.SetHeader("To", "977564830@qq.com")           //收件人
	m.SetHeader("Subject", fmt.Sprintf("打卡通知:%s",time.Now().Format("2006-01-02 15:04:05")))                     //邮件标题
	m.SetBody("text/html",content)     //邮件内容
	m.Attach(imageUrl)       //邮件附件

	d := gomail.NewDialer("smtp.163.com", 465, "rentmaterial@163.com", "UWDHGKJKRHVPYWSA")
	//邮件发送服务器信息,使用授权码而非密码
	if err = d.DialAndSend(m); err != nil {
		return
	}

	return
}

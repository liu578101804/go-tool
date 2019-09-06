package mailman

import (
	"net/smtp"
	"strings"
	"fmt"
	"strconv"
)


type Email struct {
	ServerAddr 		string
	Port 			int
	FromUser 		string
	AuthPassword 	string
	FromNickName 	string
}

func (e Email)SendEmail(to []string, subject, body string, contentType MailContentType) error  {
	//认证
	auth := smtp.PlainAuth("",e.FromUser,e.AuthPassword,e.ServerAddr)
	//组建发送体
	msg := []byte(fmt.Sprintf("To:%s\r\nFrom:%s<%s>\r\nSubject:%s\r\n%s\r\n\r\n%s",strings.Join(to,","),e.FromNickName,e.FromUser,subject,contentType,body))
	//发送
	return smtp.SendMail(
		e.ServerAddr+":"+strconv.Itoa(e.Port),
		auth,
		e.FromUser,
		to,
		msg,
	)
}

/*
QQ:
- serverAddr: smtp.qq.com
*/
func NewEmail(serverAddr,fromUser,authPassword,fromNickName string,port int) IEmail {
	return Email{
		AuthPassword: authPassword,
		ServerAddr: serverAddr,
		FromNickName: fromNickName,
		FromUser: fromUser,
		Port: port,
	}
}
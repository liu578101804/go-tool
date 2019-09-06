package mailman

import (
	"testing"
	"fmt"
)

func TestSendEmail(t *testing.T) {

	fmt.Println(NewEmail("smtp.qq.com","578101804@qq.com","smuiscrnwnxqbeaj","liu",25).SendEmail([]string{"1196990763@qq.com"},"hello","test email", MailContentTypeTxt))

}
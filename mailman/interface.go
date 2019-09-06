package mailman


type MailContentType string
const (
	MailContentTypeHtml = "Content-Type:text/html; charset=UTF-8"
	MailContentTypeTxt = "Content-Type:text/plain; charset=UTF-8"
)

type IMailman interface {
	SendEmail(to []string, subject, body string, contentType MailContentType) error
}
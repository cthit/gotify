package gotify

type MailService interface {
	SendMail(mail Mail) (Mail, error) // Returns the actually sent email
	Destroy() error
}

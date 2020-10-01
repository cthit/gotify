package mail

type Service interface {
	SendMail(mail Mail) (Mail, error) // Returns the actually sent email
	Destroy() error
}

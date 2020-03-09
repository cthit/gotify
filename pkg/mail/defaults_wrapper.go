package mail

import "strings"

type DefaultsWrapper struct {
	ms                        Service
	mailDefaultFromAddress    string
	mailDefaultReplyToAddress string
}

func NewService(ms Service, mailDefaultFromAddress string, mailDefaultReplyToAddress string) *DefaultsWrapper {
	return &DefaultsWrapper{
		ms:                        ms,
		mailDefaultFromAddress:    mailDefaultFromAddress,
		mailDefaultReplyToAddress: mailDefaultReplyToAddress,
	}
}

func (w DefaultsWrapper) SendMail(mail Mail) (Mail, error) {
	if strings.TrimSpace(mail.From) == "" {
		mail.From = w.mailDefaultFromAddress
	}

	if strings.TrimSpace(mail.ReplyTo) == "" {
		mail.ReplyTo = w.mailDefaultReplyToAddress
	}
	return w.ms.SendMail(mail)
}

func (w DefaultsWrapper) Destroy() error {
	return w.ms.Destroy()
}

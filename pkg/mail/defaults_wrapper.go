package mail

import "strings"

type DefaultsWrapper struct {
	ms                        Service
	mailDefaultFromAddress    string
	mailDefaultReplyToAddress string
	mailDefaultContentType    string
}

func NewService(ms Service, mailDefaultFromAddress, mailDefaultReplyToAddress, mailDefaultContentType string) *DefaultsWrapper {
	return &DefaultsWrapper{
		ms:                        ms,
		mailDefaultFromAddress:    mailDefaultFromAddress,
		mailDefaultReplyToAddress: mailDefaultReplyToAddress,
		mailDefaultContentType:    mailDefaultContentType,
	}
}

func (w DefaultsWrapper) SendMail(mail Mail) (Mail, error) {
	if strings.TrimSpace(mail.From) == "" {
		mail.From = w.mailDefaultFromAddress
	}

	if strings.TrimSpace(mail.ReplyTo) == "" {
		mail.ReplyTo = w.mailDefaultReplyToAddress
	}

	if strings.TrimSpace(mail.ContentType) == "" {
		mail.ContentType = w.mailDefaultContentType
	}

	return w.ms.SendMail(mail)
}

func (w DefaultsWrapper) Destroy() error {
	return w.ms.Destroy()
}

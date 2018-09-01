package mock

import (
	"github.com/cthit/gotify"
	"fmt"
)

type mockService struct {
}

func NewMockServiceCreator() (func() gotify.MailService, error) {
	return func() gotify.MailService {
		return &mockService{}
	}, nil
}

func (g *mockService) SendMail(mail gotify.Mail) (gotify.Mail, error) {

	fmt.Printf("Sending mail:\n %#v \n", mail)

	return mail, nil
}

func (g *mockService) Destroy() error {
	return nil
}
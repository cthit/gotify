package mock

import (
	"fmt"
	"github.com/cthit/gotify/pkg/mail"
)

type mockService struct {
}

func NewService() (mail.Service, error) {
	return &mockService{}, nil
}

func (g *mockService) SendMail(mail mail.Mail) (mail.Mail, error) {

	fmt.Printf("Sending mail:\n %#v \n", mail)

	return mail, nil
}

func (g *mockService) Destroy() error {
	return nil
}

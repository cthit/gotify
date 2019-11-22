package gmail

import (
	"github.com/cthit/gotify/pkg/mail"

	"google.golang.org/api/gmail/v1"

	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"

	"encoding/base64"
	"io/ioutil"
)

type googleService struct {
	mailService *gmail.Service
	adminMail   string
	debug       bool
}

func NewService(keyPath string, adminMail string, debug bool) (mail.MailService, error) {

	jsonKey, err := ioutil.ReadFile(keyPath)
	if err != nil {
		return nil, err
	}

	// Parse jsonKey
	config, err := google.JWTConfigFromJSON(jsonKey, gmail.GmailSendScope)
	if err != nil {
		return nil, err
	}

	// Why do I need this??
	config.Subject = adminMail

	mailService, err := gmail.NewService(context.Background())
	if err != nil {
		return nil, err
	}

	gs := &googleService{
		mailService: mailService,
		adminMail:   adminMail,
		debug:       debug,
	}

	return gs, err
}

func (g *googleService) SendMail(mail mail.Mail) (mail.Mail, error) {

	mail.From = g.adminMail

	msgRaw := "From: " + mail.From + "\r\n" +
		"To: " + mail.To + "\r\n" +
		"Subject: " + mail.Subject + "\r\n\r\n" +
		mail.Body + "\r\n"

	msg := &gmail.Message{
		Raw: base64.RawURLEncoding.EncodeToString([]byte(msgRaw)),
	}
	_, err := g.mailService.Users.Messages.Send(mail.From, msg).Do()

	return mail, err
}

func (g *googleService) Destroy() error {
	return nil
}

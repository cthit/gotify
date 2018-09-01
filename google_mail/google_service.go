package google_mail

import (
	"google.golang.org/api/gmail/v1" // Imports as gmail

	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"

	"encoding/base64"
	"io/ioutil"

	"github.com/cthit/gotify"
)

type googleService struct {
	mailService *gmail.Service
	adminMail   string
	debug       bool
}

func NewGoogleMailServiceCreator(keyPath string, adminMail string, debug bool) (func() gotify.MailService, error) {

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

	// Create a http client
	client := config.Client(context.Background())

	mailService, err := gmail.New(client)
	if err != nil {
		return nil, err
	}

	gs := &googleService{
		mailService: mailService,
		adminMail:   adminMail,
		debug:       debug,
	}
	if err != nil {
		return nil, err
	}

	return func() gotify.MailService {
		return gs
	}, nil
}

func (g *googleService) SendMail(mail gotify.Mail) (gotify.Mail, error) {

	mail.From = g.adminMail

	msgRaw := "From: " + mail.From + "\r\n" +
		"To: " + mail.To + "\r\n" +
		"Subject: " + mail.Subject + "\r\n\r\n" +
		mail.Body + "\r\n"

	msg := &gmail.Message{
		Raw: base64.StdEncoding.EncodeToString([]byte(msgRaw)),
	}
	_, err := g.mailService.Users.Messages.Send(mail.From, msg).Do()

	return mail, err
}

func (g *googleService) Destroy() error {
	return nil
}

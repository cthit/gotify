package gmail

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/cthit/gotify/pkg/mail"
	"golang.org/x/oauth2/google"
	"golang.org/x/oauth2/jwt"
	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/option"

	"encoding/base64"

	"golang.org/x/net/context"
)

const googleInvalidEmailErrorMessage = `Response: {
  "error": "invalid_grant",
  "error_description": "Invalid email or User ID"
}`

type googleService struct {
	config jwt.Config
	debug  bool
}

func (g *googleService) mailService(from string) (*gmail.Service, error) {
	// make sure to not edit the original config
	c := g.config
	c.Subject = from
	return gmail.NewService(context.Background(), option.WithScopes(gmail.GmailSendScope), option.WithTokenSource(c.TokenSource(context.TODO())))
}

func NewService(keyPath string, debug bool) (mail.Service, error) {
	jsonKey, err := ioutil.ReadFile(keyPath)
	if err != nil {
		return nil, err
	}

	// Parse jsonKey
	config, err := google.JWTConfigFromJSON(jsonKey, gmail.GmailSendScope)
	if err != nil {
		return nil, err
	}

	gs := &googleService{
		config: *config,
		debug:  debug,
	}

	return gs, err
}

func (g *googleService) SendMail(m mail.Mail) (mail.Mail, error) {

	mailService, err := g.mailService(m.From)
	if err != nil {
		return m, err
	}

	msgRaw := "From: " + m.From + "\r\n" +
		"To: " + m.To + "\r\n" +
		"Reply-To: " + m.ReplyTo + "\r\n" +
		"Content-Type: " + m.ContentType + "\r\n" +
		"Subject: " + mail.EncodeHeader(m.Subject) + "\r\n\r\n" +
		m.Body + "\r\n"

	msg := &gmail.Message{
		Raw: base64.RawURLEncoding.EncodeToString([]byte(msgRaw)),
	}
	_, err = mailService.Users.Messages.Send(m.From, msg).Context(context.Background()).Do()
	if err != nil {
		if strings.Contains(err.Error(), googleInvalidEmailErrorMessage) {
			return m, fmt.Errorf("Invalid from email, email must exists")
		}
		return m, err
	}

	return m, nil
}

func (g *googleService) Destroy() error {
	return nil
}

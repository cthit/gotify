package google_mail

import (
	"fmt"

	"google.golang.org/api/gmail/v1" // Imports as gmail
	"google.golang.org/api/option"

	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"

	"encoding/base64"
	"os"

	"github.com/cthit/gotify"

	"math/rand"
)

type googleService struct {
	mailService *gmail.Service
	adminMail   string
	debug       bool
}

func NewGoogleMailServiceCreator(keyPath string, adminMail string, debug bool) (func() gotify.MailService, error) {

	jsonKey, err := os.ReadFile(keyPath)
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

	mailService, err := gmail.NewService(context.Background(), option.WithHTTPClient(client))
	if err != nil {
		return nil, err
	}

	gs := &googleService{
		mailService: mailService,
		adminMail:   adminMail,
		debug:       debug,
	}

	return func() gotify.MailService {
		return gs
	}, nil
}

func (g *googleService) SendMail(mail gotify.Mail) (gotify.Mail, error) {

	mail.From = g.adminMail
	var msgRaw string
	subject := "=?UTF-8?B?" + base64.StdEncoding.EncodeToString([]byte(mail.Subject)) + "?="

	if len(mail.Attachments) > 0 {
		boundary := fmt.Sprint("gotify-boundary-", rand.Int63())

		msgRaw = "From: " + mail.From + "\r\n" +
			"To: " + mail.To + "\r\n" +
			"Subject: " + subject + "\r\n" +
			"MIME-Version: 1.0\r\n" +
			"Content-Type: multipart/mixed; boundary=" + boundary + "\r\n\r\n" +
			"--" + boundary + "\r\n" +
			"Content-Type: text/plain; charset=UTF-8\r\n\r\n" +
			mail.Body + "\r\n"

		for _, attachment := range mail.Attachments {
			msgRaw += "--" + boundary + "\r\n" +
				"Content-Type: " + attachment.ContentType + "; name=\"" + attachment.Name + "\"\r\n" +
				"Content-Disposition: attachment; filename=\"" + attachment.Name + "\"\r\n" +
				"Content-Transfer-Encoding: base64\r\n\r\n" +
				attachment.Data + "\r\n"
		}

		msgRaw += "--" + boundary + "--"
	} else {
		msgRaw = "From: " + mail.From + "\r\n" +
			"To: " + mail.To + "\r\n" +
			"Subject: " + subject + "\r\n\r\n" +
			mail.Body + "\r\n"
	}

	msg := &gmail.Message{
		Raw: base64.RawURLEncoding.EncodeToString([]byte(msgRaw)),
	}
	_, err := g.mailService.Users.Messages.Send(mail.From, msg).Do()

	return mail, err
}

func (g *googleService) Destroy() error {
	return nil
}

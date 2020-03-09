package grpc

import (
	"context"
	"strings"

	gotify "github.com/cthit/gotify/pkg/api/v1"
	"github.com/cthit/gotify/pkg/mail"
)

func (s *Server) SendMail(_ context.Context, in *gotify.Mail) (*gotify.Mail, error) {
	//validate yo
	m := mail.Mail{
		To:      in.To,
		Subject: in.Subject,
		Body:    in.Body,
	}

	if strings.TrimSpace(in.From) == "" {
		m.From = s.mailDefaultFromAddress
	} else {
		m.From = in.From
	}

	if strings.TrimSpace(in.ReplyTo) == "" {
		m.ReplyTo = s.mailDefaultReplyToAddress
	} else {
		m.ReplyTo = in.ReplyTo
	}

	m, err := s.mailService.SendMail(m)
	if err != nil {
		return nil, err
	}
	//handle them errors
	return &gotify.Mail{
		To:      m.To,
		From:    m.From,
		ReplyTo: m.ReplyTo,
		Subject: m.Subject,
		Body:    m.Body,
	}, nil
}

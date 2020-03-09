package grpc

import (
	"context"
	gotify "github.com/cthit/gotify/pkg/api/v1"
	"github.com/cthit/gotify/pkg/mail"
)

func (s *Server) SendMail(_ context.Context, in *gotify.Mail) (*gotify.Mail, error) {
	//validate yo
	m := mail.Mail{
		To:      in.To,
		From:    in.From,
		ReplyTo: in.ReplyTo,
		Subject: in.Subject,
		Body:    in.Body,
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

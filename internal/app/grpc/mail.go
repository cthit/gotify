package grpc

import (
	"context"
	gotify "github.com/cthit/gotify/pkg/api/v1"
	"github.com/cthit/gotify/pkg/mail"
)

func (s *Server) SendMail(_ context.Context, in *gotify.Mail) (*gotify.Mail, error) {
	//TODO: validate
	m := mail.Mail{
		To:      in.To,
		From:    in.From,
		ReplyTo: in.ReplyTo,
		Subject: in.Subject,
		ContentType: in.ContentType,
		Body:    in.Body,
	}

	m, err := s.mailService.SendMail(m)
	if err != nil {
		return nil, err
	}
	//TODO: handle errors
	return &gotify.Mail{
		To:      m.To,
		From:    m.From,
		ReplyTo: m.ReplyTo,
		Subject: m.Subject,
		ContentType: m.ContentType,
		Body:    m.Body,
	}, nil
}

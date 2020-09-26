package grpc

import (
	"context"

	"google.golang.org/grpc/codes"

	gotify "github.com/cthit/gotify/pkg/api/v1"
	"github.com/cthit/gotify/pkg/mail"

	"github.com/pkg/errors"
)

func (s *Server) SendMail(_ context.Context, in *gotify.Mail) (*gotify.Mail, error) {
	m := mail.Mail{
		To:          in.To,
		From:        in.From,
		ReplyTo:     in.ReplyTo,
		Subject:     in.Subject,
		ContentType: in.ContentType,
		Body:        in.Body,
	}

	err := mail.Validate(m)
	if err != nil {
		return nil, WithErrorStatus(err, codes.InvalidArgument)
	}

	m, err = s.mailService.SendMail(m)
	if err != nil {
		return nil, WithErrorStatus(errors.Wrap(err, "failed to send mail"), codes.Internal)
	}

	return &gotify.Mail{
		To:          m.To,
		From:        m.From,
		ReplyTo:     m.ReplyTo,
		Subject:     m.Subject,
		ContentType: m.ContentType,
		Body:        m.Body,
	}, nil
}

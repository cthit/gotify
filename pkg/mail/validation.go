package mail

import (
	"github.com/pkg/errors"

	"github.com/cthit/gotify/internal/validation"
)

func Validate(mail Mail) error {
	err := validation.And(
		validation.FieldString(
			"to",
			mail.To,
			validation.IsEmail,
		),
		validation.FieldString(
			"from",
			mail.From,
			validation.OrString(
				validation.IsEmail,
				validation.IsEmpty,
			),
		),
		validation.FieldString(
			"reply_to",
			mail.ReplyTo,
			validation.OrString(
				validation.IsEmail,
				validation.IsEmpty,
			),
		),
		validation.FieldString(
			"subject",
			mail.Subject,
			validation.IsNotEmpty,
		),
		validation.FieldString(
			"content_type",
			mail.ContentType,
		),
		validation.FieldString(
			"body",
			mail.Body,
			validation.IsNotEmpty,
		),
	)
	if err != nil {
		return errors.Wrap(err, "validation failed")
	}

	return nil
}

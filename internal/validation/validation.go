package validation

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
)

func Or(errs ...error) error {
	var messages []string

	for _, err := range errs {
		if err != nil {
			messages = append(messages, err.Error())
		} else {
			return nil
		}
	}

	return fmt.Errorf("should satisfy at least one of (%s)", strings.Join(messages, ", "))
}

func And(errs ...error) error {
	var messages []string

	for _, err := range errs {
		if err != nil {
			messages = append(messages, err.Error())
		}
	}

	if len(messages) == 0 {
		return nil
	} else if len(messages) == 1 {
		return errors.New(messages[0])
	}

	return fmt.Errorf("should satisfy all of (%s)", strings.Join(messages, ", "))
}

func Field(name string, errs ...error) error {
	err := And(errs...)
	if err == nil {
		return nil
	}

	return errors.Wrapf(err, "field '%s' failed validation", name)
}
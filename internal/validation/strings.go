package validation

import (
	"errors"
	"regexp"
)

var emailRegexp = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

func OrString(errFuncs ...func(string) error) func(string) error {
	return func(s string) error {
		errs := make([]error, len(errFuncs))
		for i, errFunc := range errFuncs {
			errs[i] = errFunc(s)
		}

		return Or(errs...)
	}
}

func AndString(errFuncs ...func(string) error) func(string) error {
	return func(s string) error {
		errs := make([]error, len(errFuncs))
		for i, errFunc := range errFuncs {
			errs[i] = errFunc(s)
		}

		return And(errs...)
	}
}

func FieldString(name string, value string, errFuncs ...func(string) error) error {
	errs := make([]error, len(errFuncs))
	for i, errFunc := range errFuncs {
		errs[i] = errFunc(value)
	}

	return Field(name, errs...)
}

func IsEmpty(s string) error {
	if s == "" {
		return nil
	}

	return errors.New("should be empty")
}

func IsNotEmpty(s string) error {
	if s != "" {
		return nil
	}

	return errors.New("should not be empty")
}

func IsEmail(s string) error {
	if emailRegexp.MatchString(s) {
		return nil
	}

	return errors.New("should be an email address")
}
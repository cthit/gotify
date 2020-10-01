package validation

import (
	"testing"

	"github.com/pkg/errors"
)

func willFail(s string) error {
	return errors.New("always fails")
}
func willPass(s string) error {
	return nil
}

var stringAndTests = []struct {
	name     string
	errFuncs []func(string) error
	wantErr  bool
}{
	{
		name: "no ok",
		errFuncs: []func(string) error{
			willFail,
			willFail,
			willFail,
		},
		wantErr: true,
	},
	{
		name: "fail first",
		errFuncs: []func(string) error{
			willFail,
			willPass,
			willPass,
		},
		wantErr: true,
	},
	{
		name: "fail last",
		errFuncs: []func(string) error{
			willPass,
			willPass,
			willFail,
		},
		wantErr: true,
	},
	{
		name: "fail middle",
		errFuncs: []func(string) error{
			willPass,
			willFail,
			willPass,
		},
		wantErr: true,
	},
	{
		name: "two ok ",
		errFuncs: []func(string) error{
			willFail,
			willPass,
			willPass,
		},
		wantErr: true,
	},
	{
		name: "all ok ",
		errFuncs: []func(string) error{
			willPass,
			willPass,
			willPass,
		},
		wantErr: false,
	},
	{
		name: "one fail ",
		errFuncs: []func(string) error{
			willFail,
		},
		wantErr: true,
	},
	{
		name: "one ok ",
		errFuncs: []func(string) error{
			willPass,
		},
		wantErr: false,
	},
	{
		name:     "none",
		errFuncs: []func(string) error{},
		wantErr:  false,
	},
}

func TestAndString(t *testing.T) {
	for _, tt := range stringAndTests {
		t.Run(tt.name, func(t *testing.T) {
			if err := AndString(tt.errFuncs...)(""); (err != nil) != tt.wantErr {
				t.Errorf("AndString()() = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOrString(t *testing.T) {
	tests := []struct {
		name     string
		errFuncs []func(string) error
		wantErr  bool
	}{
		{
			name: "no ok",
			errFuncs: []func(string) error{
				willFail,
				willFail,
				willFail,
			},
			wantErr: true,
		},
		{
			name: "ok first",
			errFuncs: []func(string) error{
				willPass,
				willFail,
				willFail,
			},
			wantErr: false,
		},
		{
			name: "ok last",
			errFuncs: []func(string) error{
				willFail,
				willFail,
				willPass,
			},
			wantErr: false,
		},
		{
			name: "ok middle",
			errFuncs: []func(string) error{
				willFail,
				willPass,
				willFail,
			},
			wantErr: false,
		},
		{
			name: "two ok ",
			errFuncs: []func(string) error{
				willFail,
				willPass,
				willPass,
			},
			wantErr: false,
		},
		{
			name: "all ok ",
			errFuncs: []func(string) error{
				willPass,
				willPass,
				willPass,
			},
			wantErr: false,
		},
		{
			name: "one fail ",
			errFuncs: []func(string) error{
				willFail,
			},
			wantErr: true,
		},
		{
			name: "one ok ",
			errFuncs: []func(string) error{
				willPass,
			},
			wantErr: false,
		},
		{
			name:     "none",
			errFuncs: []func(string) error{},
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := OrString(tt.errFuncs...)(""); (err != nil) != tt.wantErr {
				t.Errorf("OrString()() = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFieldString(t *testing.T) {
	for _, tt := range stringAndTests {
		t.Run(tt.name, func(t *testing.T) {
			if err := FieldString("test", "value", tt.errFuncs...); (err != nil) != tt.wantErr {
				t.Errorf("FieldString() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestIsEmail(t *testing.T) {
	tests := []struct {
		name    string
		s       string
		wantErr bool
	}{
		{
			name:    "none",
			s:       "",
			wantErr: true,
		},
		{
			name:    "spaces",
			s:       "  ",
			wantErr: true,
		},
		{
			name:    "no at",
			s:       "asdjhfb.com",
			wantErr: true,
		},
		{
			name:    "no dot",
			s:       "aaa@asdjhfb,com",
			wantErr: true,
		},
		{
			name:    "no beginning",
			s:       "@asdjhfb.com",
			wantErr: true,
		},
		{
			name:    "no end",
			s:       "aaa@asdjhfb.",
			wantErr: true,
		},
		{
			name:    "no middle",
			s:       "aaa@.asdjhfb",
			wantErr: true,
		},
		{
			name:    "valid",
			s:       "aaa@asdjhfb.com",
			wantErr: false,
		},
		{
			name:    "valid with dots",
			s:       "aa.aa.bb.cc@as.dj.h.fb.com",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := IsEmail(tt.s); (err != nil) != tt.wantErr {
				t.Errorf("IsEmail() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestIsEmpty(t *testing.T) {
	tests := []struct {
		name    string
		s       string
		wantErr bool
	}{
		{
			name:    "empty",
			s:       "",
			wantErr: false,
		},
		{
			name:    "not empty",
			s:       "a",
			wantErr: true,
		},
		{
			name:    "spaces",
			s:       "  ",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := IsEmpty(tt.s); (err != nil) != tt.wantErr {
				t.Errorf("IsEmpty() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestIsNotEmpty(t *testing.T) {
	tests := []struct {
		name    string
		s       string
		wantErr bool
	}{
		{
			name:    "empty",
			s:       "",
			wantErr: true,
		},
		{
			name:    "not empty",
			s:       "a",
			wantErr: false,
		},
		{
			name:    "spaces",
			s:       "  ",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := IsNotEmpty(tt.s); (err != nil) != tt.wantErr {
				t.Errorf("IsNotEmpty() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
package validation

import (
	"testing"

	"github.com/pkg/errors"
)

var e = errors.New("")

var andTests = []struct {
	name    string
	errs    []error
	wantErr bool
}{
	{
		name: "no ok",
		errs: []error{
			e,
			e,
			e,
		},
		wantErr: true,
	},
	{
		name: "fail first",
		errs: []error{
			e,
			nil,
			nil,
		},
		wantErr: true,
	},
	{
		name: "fail last",
		errs: []error{
			nil,
			nil,
			e,
		},
		wantErr: true,
	},
	{
		name: "fail middle",
		errs: []error{
			nil,
			e,
			nil,
		},
		wantErr: true,
	},
	{
		name: "two ok ",
		errs: []error{
			e,
			nil,
			nil,
		},
		wantErr: true,
	},
	{
		name: "all ok ",
		errs: []error{
			nil,
			nil,
			nil,
		},
		wantErr: false,
	},
	{
		name: "one fail ",
		errs: []error{
			e,
		},
		wantErr: true,
	},
	{
		name: "one ok ",
		errs: []error{
			nil,
		},
		wantErr: false,
	},
	{
		name:    "none",
		errs:    []error{},
		wantErr: false,
	},
}

func TestAnd(t *testing.T) {
	for _, tt := range andTests {
		t.Run(tt.name, func(t *testing.T) {
			if err := And(tt.errs...); (err != nil) != tt.wantErr {
				t.Errorf("And() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestField(t *testing.T) {
	for _, tt := range andTests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Field("", tt.errs...); (err != nil) != tt.wantErr {
				t.Errorf("Field() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOr(t *testing.T) {
	tests := []struct {
		name    string
		errs    []error
		wantErr bool
	}{
		{
			name: "no ok",
			errs: []error{
				e,
				e,
				e,
			},
			wantErr: true,
		},
		{
			name: "ok first",
			errs: []error{
				nil,
				e,
				e,
			},
			wantErr: false,
		},
		{
			name: "ok last",
			errs: []error{
				e,
				e,
				nil,
			},
			wantErr: false,
		},
		{
			name: "ok middle",
			errs: []error{
				e,
				nil,
				e,
			},
			wantErr: false,
		},
		{
			name: "two ok ",
			errs: []error{
				e,
				nil,
				nil,
			},
			wantErr: false,
		},
		{
			name: "all ok ",
			errs: []error{
				nil,
				nil,
				nil,
			},
			wantErr: false,
		},
		{
			name: "one fail ",
			errs: []error{
				e,
			},
			wantErr: true,
		},
		{
			name: "one ok ",
			errs: []error{
				nil,
			},
			wantErr: false,
		},
		{
			name:    "none",
			errs:    []error{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Or(tt.errs...); (err != nil) != tt.wantErr {
				t.Errorf("Or() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
package mail

import "testing"

func TestValidate(t *testing.T) {
	tests := []struct {
		name    string
		mail    Mail
		wantErr bool
	}{
		{
			name:    "empty mail",
			mail:    Mail{},
			wantErr: true,
		},
		{
			name: "missing to",
			mail: Mail{
				Subject: "not empty",
				Body:    "not empty",
			},
			wantErr: true,
		},
		{
			name: "invalid to",
			mail: Mail{
				To:      "qwerty@abgc,com",
				Subject: "not empty",
				Body:    "not empty",
			},
			wantErr: true,
		},
		{
			name: "no optional fields",
			mail: Mail{
				To:      "qwerty@abgc.com",
				Subject: "not empty",
				Body:    "not empty",
			},
			wantErr: false,
		},
		{
			name: "no subject",
			mail: Mail{
				To:   "qwerty@abgc.com",
				Body: "not empty",
			},
			wantErr: true,
		},
		{
			name: "no body",
			mail: Mail{
				To:      "qwerty@abgc.com",
				Subject: "not empty",
			},
			wantErr: true,
		},
		{
			name: "invalid from",
			mail: Mail{
				From:    "qwerty@abgc,com",
				To:      "qwerty@abgc.com",
				Subject: "not empty",
				Body:    "not empty",
			},
			wantErr: true,
		},
		{
			name: "invalid reply-to",
			mail: Mail{
				ReplyTo: "qwerty@abgc,com",
				To:      "qwerty@abgc.com",
				Subject: "not empty",
				Body:    "not empty",
			},
			wantErr: true,
		},
		{
			name: "all fields",
			mail: Mail{
				To:          "qwerty@abgc.com",
				From:        "qwerty@abgc.com",
				ReplyTo:     "qwerty@abgc.com",
				Subject:     "not empty",
				ContentType: "some content type",
				Body:        "not empty",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Validate(tt.mail); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
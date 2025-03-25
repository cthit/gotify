package gotify

type Mail struct {
	To          string       `json:"to"`
	From        string       `json:"from"`
	Subject     string       `json:"subject"`
	Body        string       `json:"body"`
	Attachments []Attachment `json:"attachments"`
}

type Attachment struct {
	Name        string `json:"name"`
	Data        string `json:"data"`
	ContentType string `json:"content_type"`
}

package mail

type Mail struct {
	To          string `json:"to"`
	From        string `json:"from"`
	ReplyTo     string `json:"reply_to"`
	Subject     string `json:"subject"`
	ContentType string `json:"content_type"`
	Body        string `json:"body"`
}

package mail

type Mail struct {
	To      string `json:"to"`
	From    string `json:"from"`
	ReplyTo string `json:"reply:to"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
}

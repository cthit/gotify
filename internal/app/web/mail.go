package web

import (
	"encoding/json"
	"github.com/cthit/gotify/pkg/mail"
	"github.com/gocraft/web"
	"io/ioutil"
	"net/http"
)

func (c *Context) SendMail(rw web.ResponseWriter, req *web.Request) {
	var m mail.Mail

	// Read request body
	body, err := ioutil.ReadAll(req.Body)
	req.Body.Close()
	if err != nil {
		c.printError(err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	// Parse json email
	err = json.Unmarshal(body, &m)
	if err != nil {
		c.printError(err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	// Send email
	m, err = c.MailService.SendMail(m)
	if err != nil {
		c.printError(err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Build json email
	data, err := json.Marshal(m)
	if err != nil {
		c.printError(err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Return the sent email
	rw.WriteHeader(http.StatusOK)
	rw.Write(data)
}

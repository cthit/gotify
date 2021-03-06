package web

import (
	"encoding/json"
	"github.com/cthit/gotify"
	"github.com/gocraft/web"
	"io/ioutil"
	"net/http"
)

func (c *Context) SendMail(rw web.ResponseWriter, req *web.Request) {
	var mail gotify.Mail

	// Read request body
	body, err := ioutil.ReadAll(req.Body)
	req.Body.Close()
	if err != nil {
		c.printError(err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	// Parse json email
	err = json.Unmarshal(body, &mail)
	if err != nil {
		c.printError(err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	// Send email
	mail, err = c.MailService.SendMail(mail)
	if err != nil {
		c.printError(err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Build json email
	data, err := json.Marshal(mail)
	if err != nil {
		c.printError(err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Return the sent email
	rw.WriteHeader(http.StatusOK)
	rw.Write(data)
}

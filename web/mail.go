package web

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/cthit/gotify"
	"github.com/gocraft/web"
	"github.com/spf13/viper"
)

func (c *Context) SendMail(rw web.ResponseWriter, req *web.Request) {
	var mail gotify.Mail

	// Ensure that the request is not too large
	if req.ContentLength > viper.GetInt64("max-mail-size") {
		rw.WriteHeader(http.StatusRequestEntityTooLarge)
		return
	}

	// Read request body
	body, err := io.ReadAll(req.Body)
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

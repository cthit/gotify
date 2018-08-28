package web

import (
	"github.com/gocraft/web"
	"io/ioutil"
	"fmt"
	"net/http"
	"encoding/json"
	"../../gotify"
)

func (c *Context) SendMail(rw web.ResponseWriter, req *web.Request) {
	var mail gotify.Mail

	// Read request body
	body, err := ioutil.ReadAll(req.Body)
	req.Body.Close()
	if err != nil {
		fmt.Println(err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	// Parse json email
	err = json.Unmarshal(body, &mail)
	if err != nil {
		fmt.Println(err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	// Send email
	mail, err = c.MailService.SendMail(mail)
	if err != nil {
		fmt.Println(err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Build json email
	data, err := json.Marshal(mail)
	if err != nil {
		fmt.Println(err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Return the sent email
	rw.WriteHeader(http.StatusOK)
	rw.Write(data)
}

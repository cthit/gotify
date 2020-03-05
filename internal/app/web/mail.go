package web

import (
	"fmt"
	"github.com/cthit/gotify/pkg/mail"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *Server) MailHandler(c *gin.Context) {
	var m mail.Mail

	err := c.Bind(&m)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Send email
	m, err = s.mailService.SendMail(m)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	// Return the sent email
	c.JSON(http.StatusAccepted, m)
}

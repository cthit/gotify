package app

import (
	"github.com/cthit/gotify/internal/app/config"
	"github.com/cthit/gotify/internal/app/web"
	"github.com/cthit/gotify/pkg/mail"
	"github.com/cthit/gotify/pkg/mail/gmail"
	"github.com/cthit/gotify/pkg/mail/mock"

	"fmt"
	"log"
	"net/http"
)


func Start() error {
	c, err := config.LoadConfig()
	if err != nil {
		fmt.Println("Failed to load config.")
		return err
	} else {
		fmt.Println("Loaded config.")
	}

	fmt.Printf("Debug mode is set to: %t \n", c.Debug())
	fmt.Printf("Mock mode is set to: %t \n", c.Mock())

	fmt.Printf("Setting up services...")

	var mailService mail.MailService

	if !c.Mock() {
		mailService, err = gmail.NewService(
			c.GmailKeyfile(),
			c.GmailAdminMail(),
			c.Debug(),
		)
		if err != nil {
			return err
		}
	} else {
		mailService, _ = mock.NewService()
	}

	preSharedKey := c.PreSharedKey()

	fmt.Printf("Done! \n")

	fmt.Printf("Serving application on port %s \n", c.Port())
	log.Fatal(
		http.ListenAndServe(":"+c.Port(),
			web.Router(
				preSharedKey,
				mailService,
				c.Debug(),
			),
		),
	)
	return nil
}

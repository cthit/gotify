package app

import (
	"github.com/cthit/gotify/internal/app/config"
	"github.com/cthit/gotify/internal/app/grpc"
	"github.com/cthit/gotify/pkg/mail"
	"github.com/cthit/gotify/pkg/mail/gmail"
	"github.com/cthit/gotify/pkg/mail/mock"

	"fmt"
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
	fmt.Printf("Environment is set to: %s \n", c.Environment())

	fmt.Printf("Setting up services...")

	var mailService mail.Service

	if !c.Mock() {
		mailService, err = gmail.NewService(
			c.GmailKeyfile(),
			c.Debug(),
		)
		if err != nil {
			return err
		}
	} else {
		mailService, _ = mock.NewService()
	}

	mailService = mail.NewService(mailService, c.MailDefaultFrom(), c.MailDefaultReplyTo(), c.MailDefaultContentType())

	fmt.Printf("Done! \n")

	fmt.Printf("Serving application on port %s \n", c.WebPort())
	fmt.Printf("Serving rpc on port %s \n", c.RPCPort())
	server, err := grpc.NewServer(
		c.RPCPort(),
		c.WebPort(),
		c.Environment(),
		c.Debug(),
		mailService,
	)
	if err != nil {
		fmt.Println("Failed to create webserver.")
		return err
	}

	server.Start()
	return nil
}
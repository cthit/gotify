package main

import (
	"github.com/cthit/gotify/google_mail"
	"github.com/cthit/gotify/web"
	"log"
	"net/http"
	"fmt"
	"github.com/spf13/viper"
	"github.com/cthit/gotify/mock"
	"github.com/cthit/gotify"
)


func init() {
	err := loadConfig()
	if err != nil {
		fmt.Println("Failed to load config.")
	} else {
		fmt.Println("Loaded config.")

	}
}

func main() {

	web.DEBUG = viper.GetBool("debug-mode")
	fmt.Printf("Debug mode is set to: %t \n", viper.GetBool("debug-mode"))
	fmt.Printf("Mock mode is set to: %t \n", viper.GetBool("mock-mode"))

	fmt.Printf("Setting up services...")

	var mailServiceCreator func()  gotify.MailService
	var err error

	if !viper.GetBool("mock-mode") {
		mailServiceCreator, err = google_mail.NewGoogleMailServiceCreator(viper.GetString("google-mail.keyfile"), viper.GetString("google-mail.admin-mail"))
		if err != nil {
			panic(err)
		}
	}else {
		mailServiceCreator, _ = mock.NewMockServiceCreator()
	}

	preSharedKey := viper.GetString("pre-shared-key")

	fmt.Printf("Done! \n")

	fmt.Printf("Serving application on port %s \n", viper.GetString("port"))
	log.Fatal(http.ListenAndServe(":" + viper.GetString("port"), web.Router(preSharedKey, mailServiceCreator)))
}

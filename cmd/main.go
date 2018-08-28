package main

import (
	"github.com/cthit/gotify/google_mail"
	"github.com/cthit/gotify/web"
	"log"
	"net/http"
	"fmt"
	"github.com/spf13/viper"
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

	fmt.Printf("Setting up services...")

	mailServiceCreator, err := google_mail.NewGoogleMailServiceCreator(viper.GetString("google-mail.keyfile"), viper.GetString("google-mail.admin-mail"))
	if err != nil {
		panic(err)
	}

	preSharedKey := viper.GetString("pre-shared-key")


	fmt.Printf("Done! \n")

	fmt.Printf("Serving application on port %s \n", viper.GetString("port"))
	log.Fatal(http.ListenAndServe(":" + viper.GetString("port"), web.Router(preSharedKey, mailServiceCreator)))
}

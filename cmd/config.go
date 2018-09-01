package main

import (
	"github.com/spf13/viper"
)

func loadConfig() error {
	viper.SetDefault("port", "8080")
	viper.SetDefault("debug-mode", false)
	viper.SetDefault("google-mail.keyfile", "gapps.json")
	viper.SetDefault("mock-mode", false)

	viper.SetEnvPrefix("gotify")
	viper.AutomaticEnv()

	viper.SetConfigName("config")         // name of config file (without extension)
	viper.AddConfigPath("/etc/gotify/")   // path to look for the config file in
	viper.AddConfigPath("$HOME/.gotify/") // call multiple times to add many search paths
	viper.AddConfigPath(".")              // optionally look for config in the working directory

	err := viper.ReadInConfig() // Find and read the config file
	return err

}

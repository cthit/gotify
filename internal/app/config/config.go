package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {}

func LoadConfig() (*Config, error) {
	viper.SetDefault("port", "8080")
	viper.SetDefault("debug-mode", false)
	viper.SetDefault("google-mail.keyfile", "gapps.json")
	viper.SetDefault("mock-mode", false)

	viper.SetEnvPrefix("gotify")
	viper.AutomaticEnv()

	viper.SetConfigName("config")
	viper.AddConfigPath("/etc/gotify/")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("Failed to read config from file")
		} else {
			return &Config{}, err
		}
	}
	return &Config{}, nil
}

func (*Config) Port() string {
	return viper.GetString("port")
}

func (*Config) Debug() bool {
	return viper.GetBool("debug-mode")
}

func (*Config) PreSharedKey() string {
	return viper.GetString("pre-shared-key")
}

func (*Config) Mock() bool {
	return viper.GetBool("mock-mode")
}

func (*Config) GmailKeyfile() string {
	return viper.GetString("google-mail.keyfile")
}
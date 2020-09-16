package config

import (
	"fmt"
	"github.com/spf13/viper"
	"strings"
)

type Config struct{}

func LoadConfig() (*Config, error) {
	viper.SetDefault("web-port", "8080")
	viper.SetDefault("rpc-port", "8090")
	viper.SetDefault("debug-mode", false)
	viper.SetDefault("google-mail.keyfile", "gapps.json")
	viper.SetDefault("mail.default-from", "admin@chalmers.it")
	viper.SetDefault("mail.default-reply-to", "no-reply@chalmers.it")
	viper.SetDefault("mail.default-content-type", "text/html; charset=ISO-8859-1")
	viper.SetDefault("mock-mode", false)

	viper.SetEnvPrefix("gotify")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_", ".", "_"))

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

func (*Config) WebPort() string {
	return viper.GetString("web-port")
}

func (*Config) Debug() bool {
	return viper.GetBool("debug-mode")
}

func (*Config) RPCPort() string {
	return viper.GetString("rpc-port")
}

func (*Config) Mock() bool {
	return viper.GetBool("mock-mode")
}

func (*Config) GmailKeyfile() string {
	return viper.GetString("google-mail.keyfile")
}

func (*Config) MailDefaultFrom() string {
	return viper.GetString("mail.default-from")
}

func (*Config) MailDefaultReplyTo() string {
	return viper.GetString("mail.default-reply-to")
}

func (*Config) MailDefaultContentType() string {
	return viper.GetString("mail.default-content-type")
}

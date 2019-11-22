package config

import (
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
	return &Config{}, err
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

func (*Config) GmailAdminMail() string {
	return viper.GetString("google-mail.admin-mail")
}
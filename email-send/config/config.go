package config

import "github.com/spf13/viper"

type Config struct {
	SmtpHost     string
	SmtpPort     string
	SenderEmail  string
	Password     string
	ListenerPort string
}

func New() (*Config, error) {
	viper.AddConfigPath(".")
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	cfg := Config{
		SmtpHost:     viper.GetString("SMTP_HOST"),
		SmtpPort:     viper.GetString("SMTP_PORT"),
		SenderEmail:  viper.GetString("SENDER_EMAIL"),
		Password:     viper.GetString("PASSWORD"),
		ListenerPort: viper.GetString("LISTENER_PORT"),
	}
	return &cfg, nil
}

package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Port     string
	JWTKey   string
	JWTTTL   int
	PgURL    string
	GRPCPort string
}

func New() (*Config, error) {
	viper.AddConfigPath(".")
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	url := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		viper.GetString("DB_HOST"),
		viper.GetString("DB_PORT"),
		viper.GetString("DB_USER"),
		viper.GetString("DB_PASSWORD"),
		viper.GetString("DB_NAME"),
		viper.GetString("DB_SSL_MODE"),
	)

	cfg := Config{
		Port:     viper.GetString("PORT"),
		JWTKey:   viper.GetString("JWT_SECRET_KEY"),
		JWTTTL:   viper.GetInt("JWT_TTL"),
		PgURL:    url,
		GRPCPort: viper.GetString("GRPC_PORT"),
	}

	return &cfg, nil
}

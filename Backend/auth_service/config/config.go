package config

import (
	flags "github.com/jessevdk/go-flags"
)

type Config struct {
	Host         string `long:"host" env:"AUTH_SERVICE_HOST"`
	Port         string `long:"port" env:"AUTH_SERVICE_PORT"`
	AuthDbName   string `long:"auth-db-name" env:"AUTH_DB_NAME"`
	AuthDbHost   string `long:"auth-db-host" env:"AUTH_DB_HOST"`
	AuthDbPort   string `long:"auth-db-port" env:"AUTH_DB_PORT"`
	AuthDbUser   string `long:"auth-db-user" env:"AUTH_DB_USER"`
	AuthDbPass   string `long:"auth-db-pass" env:"AUTH_DB_PASS"`
	JwtSecretKey string `long:"jwt-secret-key" env:"JWT_SECRET_KEY"`
}

func LoadConfig() (Config, error) {
	var cfg Config
	parser := flags.NewParser(&cfg, flags.Default)
	_, err := parser.Parse()
	return cfg, err
}

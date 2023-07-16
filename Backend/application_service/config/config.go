package config

import (
	flags "github.com/jessevdk/go-flags"
)

type Config struct {
	Host              string `long:"host" env:"APPLICATION_SERVICE_HOST"`
	Port              string `long:"port" env:"APPLICATION_SERVICE_PORT"`
	ApplicationDbName string `long:"application-db-name" env:"APPLICATION_DB_NAME"`
	ApplicationDbHost string `long:"application-db-host" env:"APPLICATION_DB_HOST"`
	ApplicationDbPort string `long:"application-db-port" env:"APPLICATION_DB_PORT"`
	ApplicationDbUser string `long:"application-db-user" env:"APPLICATION_DB_USER"`
	ApplicationDbPass string `long:"application-db-pass" env:"APPLICATION_DB_PASS"`
}

func LoadConfig() (Config, error) {
	var cfg Config
	parser := flags.NewParser(&cfg, flags.Default)
	_, err := parser.Parse()
	return cfg, err
}

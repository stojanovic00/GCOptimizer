package config

import (
	flags "github.com/jessevdk/go-flags"
)

type Config struct {
	Host        string `long:"host" env:"HOST"`
	PublicPort  string `long:"public-port" env:"PUBLIC_PORT"`
	PrivatePort string `long:"private-port" env:"PRIVATE_PORT"`
}

func LoadConfig() (Config, error) {
	var cfg Config
	parser := flags.NewParser(&cfg, flags.Default)
	_, err := parser.Parse()
	return cfg, err
}

package config

import (
	flags "github.com/jessevdk/go-flags"
)

type Config struct {
	Host                   string `long:"host" env:"API_GATEWAY_HOST"`
	PublicPort             string `long:"public-port" env:"API_GATEWAY_PUBLIC_PORT"`
	PrivatePort            string `long:"private-port" env:"API_GATEWAY_PRIVATE_PORT"`
	AuthServiceHost        string `long:"auth_service_host" env:"AUTH_SERVICE_HOST"`
	AuthServicePort        string `long:"auth_service_port" env:"AUTH_SERVICE_PORT"`
	ApplicationServiceHost string `long:"application_service_host" env:"APPLICATION_SERVICE_HOST"`
	ApplicationServicePort string `long:"application_service_port" env:"APPLICATION_SERVICE_PORT"`
	SchedulingServiceHost  string `long:"scheduling_service_host" env:"SCHEDULING_SERVICE_HOST"`
	SchedulingServicePort  string `long:"scheduling_service_port" env:"SCHEDULING_SERVICE_PORT"`
}

func LoadConfig() (Config, error) {
	var cfg Config
	parser := flags.NewParser(&cfg, flags.Default)
	_, err := parser.Parse()
	return cfg, err
}

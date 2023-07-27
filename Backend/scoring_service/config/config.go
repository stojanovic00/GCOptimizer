package config

import (
	flags "github.com/jessevdk/go-flags"
)

type Config struct {
	Host string `long:"host" env:"SCORING_SERVICE_HOST"`
	Port string `long:"port" env:"SCORING_SERVICE_PORT"`

	ScoringDbName string `long:"scoring-db-name" env:"SCORING_DB_NAME"`
	ScoringDbHost string `long:"scoring-db-host" env:"SCORING_DB_HOST"`
	ScoringDbPort string `long:"scoring-db-port" env:"SCORING_DB_PORT"`
	ScoringDbUser string `long:"scoring-db-user" env:"SCORING_DB_USER"`
	ScoringDbPass string `long:"scoring-db-pass" env:"SCORING_DB_PASS"`

	SchedulingServiceHost  string `long:"scheduling-service-host" env:"SCHEDULING_SERVICE_HOST"`
	SchedulingServicePort  string `long:"scheduling-service-port" env:"SCHEDULING_SERVICE_PORT"`
	ApplicationServiceHost string `long:"application-service-host" env:"APPLICATION_SERVICE_HOST"`
	ApplicationServicePort string `long:"application-service-port" env:"APPLICATION_SERVICE_PORT"`
	AuthServiceHost        string `long:"auth-service-host" env:"AUTH_SERVICE_HOST"`
	AuthServicePort        string `long:"auth-service-port" env:"AUTH_SERVICE_PORT"`
}

func LoadConfig() (Config, error) {
	var cfg Config
	parser := flags.NewParser(&cfg, flags.Default)
	_, err := parser.Parse()
	return cfg, err
}

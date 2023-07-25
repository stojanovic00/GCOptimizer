package config

import (
	flags "github.com/jessevdk/go-flags"
)

type Config struct {
	Host string `long:"host" env:"APPLICATION_SERVICE_HOST"`
	Port string `long:"port" env:"APPLICATION_SERVICE_PORT"`

	ScoringDbName string `long:"scoring-db-name" env:"SCORING_DB_NAME"`
	ScoringDbHost string `long:"scoring-db-host" env:"SCORING_DB_HOST"`
	ScoringDbPort string `long:"scoring-db-port" env:"SCORING_DB_PORT"`
	ScoringDbUser string `long:"scoring-db-user" env:"SCORING_DB_USER"`
	ScoringDbPass string `long:"scoring-db-pass" env:"SCORING_DB_PASS"`
}

func LoadConfig() (Config, error) {
	var cfg Config
	parser := flags.NewParser(&cfg, flags.Default)
	_, err := parser.Parse()
	return cfg, err
}

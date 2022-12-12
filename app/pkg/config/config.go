package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	HttpAddress string `default:"127.0.0.1:8000"`
	DBAddress   string `default:"host=localhost user=postgres password=1234 dbname=postgres port=5432 sslmode=disable"`
}

func NewConfig() (*Config, error) {
	var s Config
	err := envconfig.Process("myapp", &s)
	if err != nil {
		return nil, err
	}

	return &s, nil
}

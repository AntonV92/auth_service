package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	DbConf DbConfig
}

type DbConfig struct {
	User   string `env:"DB_USER"`
	Pass   string `env:"DB_PASSWORD"`
	Host   string `env:"DB_HOST"`
	Port   string `env:"DB_PORT"`
	DbName string `env:"DB_NAME"`
}

func LoadConfig() (*Config, error) {

	cfg := Config{}

	err := cleanenv.ReadConfig("config/.env", &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}

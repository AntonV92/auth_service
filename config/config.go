package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	DbConf DbConfig
}

type DbConfig struct {
	User    string `env:"DB_USER"`
	Pass    string `env:"DB_PASSWORD"`
	Host    string `env:"DB_HOST"`
	Port    string `env:"DB_PORT"`
	DbName  string `env:"DB_NAME"`
	SSLMode string `env:"DB_SSL_MODE"`
}

func LoadConfig() (*Config, error) {

	cfg := Config{}

	err := cleanenv.ReadConfig("config/.env", &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}

func (conf *Config) GetDataSource() string {
	dbConf := conf.DbConf

	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		dbConf.Host, dbConf.Port, dbConf.User, dbConf.Pass, dbConf.DbName, dbConf.SSLMode)
}

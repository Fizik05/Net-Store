package config

import (
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env        string     `yaml:"env" env:"env" env-default:"local"`
	HTTPServer HTTPServer `yaml:",inline" env:",inline"`
	Postgres   Postgres   `yaml:",inline" env:",inline"`
}

type HTTPServer struct {
	User     string `yaml:"HTTP_USER" env:"HTTP_USER"`
	Password string `yaml:"HTTP_PASSWORD" env:"HTTP_PASSWORD"`
	Address  string `yaml:"HTTP_ADDRESS" env:"HTTP_ADDRESS"`
}

type Postgres struct {
	Host     string `yaml:"POSTGRES_HOST" env:"POSTGRES_HOST"`
	Port     string `yaml:"POSTGRES_PORT" env:"POSTGRES_PORT"`
	Dbname   string `yaml:"POSTGRES_DB" env:"POSTGRES_DB"`
	Username string `yaml:"POSTGRES_USER" env:"POSTGRES_USER"`
	Password string `yaml:"POSTGRES_PASSWORD" env:"POSTGRES_PASSWORD"`
}

func MustLoad() *Config {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		log.Fatal("CONFIG_PATH is not set")
	}

	// check if file exests
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exists: %s", configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("cannot read config: %s", err)
	}

	return &cfg
}

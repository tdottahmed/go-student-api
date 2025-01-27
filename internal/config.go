package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type HTTPServer struct {
	Addr string `yaml:"address" env:"HTTP_SERVER" default:"localhost:8082" required:"true"`
}

type Config struct {
	Env        string     `yaml:"env" env:"ENV" default:"dev" required:"true"`
	Storage    string     `yaml:"storage_path" env:"STORAGE_PATH" default:"storage/storage.db" required:"true"`
	HTTPServer HTTPServer `yaml:"http_server"`
}

func MustLoadConfig() *Config {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		flag.StringVar(&configPath, "config", "config.yml", "config file path")
		flag.Parse()
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file not found: %s", configPath)
	}

	var config Config
	if err := cleanenv.ReadConfig(configPath, &config); err != nil {
		log.Fatalf("failed to load config: %s", err)
	}

	return &config
}

package internal

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type HTTPServer struct {
	Addr string `yaml:"http_server" env:"HTTP_SERVER" default:"localhost:8082" required:"true"`
}

type Config struct {
	Env     string `yaml:"env" env:"ENV" default:"dev" required:"true"`
	Storage string `yaml:"storage_path" env:"STORAGE_PATH" default:"storage/storage.db" required:"true"`
	HTTPServer
}

func MustLoadConfig() *Config {
	var configPath string
	configPath = os.Getenv("CONFIG_PATH")
	if configPath == "" {
		flags := flag.String("config", "", "config file path")
		flag.Parse()
		configPath = *flags
	}

	if configPath == "" {
		log.Fatal("config file path is required")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file not found: %s", configPath)
	}
	var config Config
	err := cleanenv.ReadConfig(configPath, &config)
	if err != nil {
		log.Fatalf("failed to load config: %s", err)
	}
	return &config
}

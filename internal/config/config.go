package config

import (
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env         string `yaml:"env" env-default:"local"`
	PostgresUrl string `yaml:"postgres_url" env-required:"true"`
	HTTPServer  `yaml:"http_server"`
}

type HTTPServer struct {
	Address     string        `yaml:"address" env-required:"true"`
	// Timeout     time.Duration `yaml:"timeout" env-required:"true"`
	// IdleTimeout time.Duration `yaml:"idle_timeout" env-required:"true"`
}

func MustLoad() (cfg Config) {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		log.Fatal("need set the CONFIG_PATH variable in .env file")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("%s does not exist, err: %s", configPath, err.Error())
	}

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("Read config %s", err.Error())
	}

	return
}

package config

import (
	"os"

	"github.com/golobby/dotenv"
)

type Config struct {
	Database struct {
		Name     string `env:"DB_NAME"`
		Host     string `env:"DB_HOST"`
		User     string `env:"DB_USER"`
		Password string `env:"DB_PASSWORD"`
	}
}

func LoadConfig() Config {
	envConfig := Config{}
	file, err := os.Open(".env")

	if err != nil {
		panic(err)
	}

	err = dotenv.NewDecoder(file).Decode(&envConfig)

	return envConfig
}

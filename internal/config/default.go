package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PostgreDriver  string `mapstructure:"POSTGRES_DRIVER"`
	PostgresSource string `mapstructure:"POSTGRES_SOURCE"`
	ServerPort     string `mapstructure:"SERVER_PORT"`
	ClientPort     string `mapstructure:"CLIENT_PORT"`
	Origin         string `mapstructure:"ORIGIN"`
}

var config Config

func LoadConfig() (err error) {

	if err := godotenv.Load(); err != nil {
		fmt.Println("Load env file err: ", err)
	}

	config = Config{
		PostgreDriver:  os.Getenv("POSTGRES_DRIVER"),
		PostgresSource: os.Getenv("POSTGRES_SOURCE"),
		ServerPort:     os.Getenv("SERVER_PORT"),
		ClientPort:     os.Getenv("CLIENT_PORT"),
	}

	return nil
}

func GetConfig() *Config {
	return &config
}

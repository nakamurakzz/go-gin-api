package config

import (
	"log"
	"os"
	"strconv"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

type Config struct {
	Env  string
	Port int
}

func New() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT")
	portInt, err := strconv.Atoi(port)
	if err != nil {
		return nil, err
	}

	cfg := &Config{
		Env:  "dev",
		Port: portInt,
	}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}

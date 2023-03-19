package config

import (
	"log"
	"os"
	"strconv"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

type Config struct {
	Env         string
	Port        int
	DB_HOST     string
	DB_PORT     string
	DB_USER     string
	DB_PASS     string
	DB_DATABASE string
}

func New() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbDatabase := os.Getenv("DB_DATABASE")
	portInt, err := strconv.Atoi(port)
	if err != nil {
		return nil, err
	}

	cfg := &Config{
		Env:         "dev",
		Port:        portInt,
		DB_HOST:     dbHost,
		DB_PORT:     dbPort,
		DB_USER:     dbUser,
		DB_PASS:     dbPass,
		DB_DATABASE: dbDatabase,
	}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}

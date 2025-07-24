package config

import (
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	BotToken string
	ChatID   string
	Feeds    []string
}

func LoadCofig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return &Config{
		BotToken: os.Getenv("TELEGRAM_BOT_TOKEN"),
		ChatID:   os.Getenv("USER_CHAT_ID"),
		Feeds:    strings.Split(os.Getenv("RSS_FEEDS"), ","),
	}
}

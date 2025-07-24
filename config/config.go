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

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file, using environment variables")
	}

	botToken := os.Getenv("TELEGRAM_BOT_TOKEN")
	chatID := os.Getenv("USER_CHAT_ID")
	rssFeeds := os.Getenv("RSS_FEEDS")

	if botToken == "" || chatID == "" || rssFeeds == "" {
		log.Fatal("Missing required environment variables. Please check your .env file or environment.")
	}

	return &Config{
		BotToken: botToken,
		ChatID:   chatID,
		Feeds:    strings.Split(rssFeeds, ","),
	}
}

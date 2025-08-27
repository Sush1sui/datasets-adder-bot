package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port     string
	ServerURL string
	BotToken string
	GuildID string
}

var Global *Config

func New() error {
	if err := godotenv.Load(); err != nil { fmt.Println("Error loading .env file") }

	port := os.Getenv("PORT")
	if port == "" {
		port = "9969"
	}
	serverUrl := os.Getenv("SERVER_URL")
	if serverUrl == "" {
		fmt.Println("SERVER_URL is not set")
	}
	botToken := os.Getenv("BOT_TOKEN")
	if botToken == "" {
		return fmt.Errorf("BOT_TOKEN is required")
	}
	guildID := os.Getenv("GUILD_ID")
	if guildID == "" {
		return fmt.Errorf("GUILD_ID is required")
	}

	Global = &Config{
		Port:     port,
		ServerURL: serverUrl,
		BotToken: botToken,
		GuildID:  guildID,
	}
	fmt.Println("Config initialized successfully")
	return nil
}
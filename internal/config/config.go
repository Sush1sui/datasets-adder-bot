package config

import (
	"context"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

type Config struct {
	Port     string
	ServerURL string
	BotToken string
	GuildID string
	MongoDBName string
	MongoDBUserAccountName string
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
	mongoDBName := os.Getenv("MONGODB_NAME")
	if mongoDBName == "" {
		return fmt.Errorf("MONGODB_NAME is required")
	}
	mongoDBUserAccountName := os.Getenv("MONGODB_USERACCOUNT_NAME")
	if mongoDBUserAccountName == "" {
		return fmt.Errorf("MONGODB_USERACCOUNT_NAME is required")
	}

	Global = &Config{
		Port:     port,
		ServerURL: serverUrl,
		BotToken: botToken,
		GuildID:  guildID,
		MongoDBName: mongoDBName,
		MongoDBUserAccountName: mongoDBUserAccountName,
	}
	fmt.Println("Config initialized successfully")
	return nil
}


func MongoConnection() *mongo.Client {
	// Use the SetServerAPIOptions() method to set the version of the Stable API on the client
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(os.Getenv("MONGO_URI")).SetServerAPIOptions(serverAPI)

	// Create a new client and connect to the server
  client, err := mongo.Connect(opts)
  if err != nil {
    panic(err)
  }

  // Send a ping to confirm a successful connection
  if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
    panic(err)
  }
  fmt.Println("DB Connected!")

	return client
}
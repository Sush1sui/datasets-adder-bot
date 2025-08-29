package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/Sush1sui/datasets_adder/internal/bot"
	"github.com/Sush1sui/datasets_adder/internal/common"
	"github.com/Sush1sui/datasets_adder/internal/config"
	"github.com/Sush1sui/datasets_adder/internal/repository"
	"github.com/Sush1sui/datasets_adder/internal/repository/mongodb"
	"github.com/Sush1sui/datasets_adder/internal/server"
)

func main() {
	err := config.New()
	if err != nil {
		fmt.Println("Error initializing configuration:", err)
	}

	mongoClient := config.MongoConnection()
	defer mongoClient.Disconnect(context.Background())
	if err := mongoClient.Ping(context.Background(), nil); err != nil {
		panic(fmt.Errorf("failed to connect to MongoDB: %w", err))
	}

	userAccountCollection := mongoClient.Database(config.Global.MongoDBName).Collection(config.Global.MongoDBUserAccountName)
	repository.UserAccountService = &mongodb.MongoClient{
		Client: userAccountCollection,
	}

	addr := fmt.Sprintf(":%s", config.Global.Port)
	router := server.NewRouter()
	fmt.Printf("Server is running on PORT: %s\n", addr)

	go func() {
		if err := http.ListenAndServe(addr, router); err != nil {
			fmt.Printf("Error starting server: %v\n", err)
		}
	}()

	go bot.StartBot()

	go func() {
		common.PingServerLoop(config.Global.ServerURL)
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	<-sigChan
	fmt.Println("Shutting down server gracefully...")
}
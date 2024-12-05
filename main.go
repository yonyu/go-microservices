package main

import (
	"github.com/yonyu/go-microservices/internal/database"
	"github.com/yonyu/go-microservices/internal/server"
	"log"
)

func main() {
	db, err := database.NewDatabaseClient()
	if err != nil {
		log.Fatalf("Failed to initialize Database Client: %s", err)
	}
	srv := server.NewEchoServer(db)
	if err := srv.Start(); err != nil {
		log.Fatalf("Failed to start server: %s", err)
	}
}

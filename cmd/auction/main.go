package main

import (
	"context"
	"github.com/zemartins81/posgo_leilao/configuration/database/mongodb"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	ctx := context.Background()
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
		return
	}

	databaseClient, err := mongodb.NewMongoDBConnections(ctx)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
}

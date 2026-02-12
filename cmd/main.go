package main

import (
	"log"
	"movie-api/internal/db"
	"movie-api/internal/handlers"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	err := db.Init()
	if err != nil {
		log.Fatalf("failed to init dynamodb: %v", err)
	}
	lambda.Start(handlers.Router)
}

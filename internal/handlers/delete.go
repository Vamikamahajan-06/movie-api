package handlers

import (
	"context"
	"log"
	"strings"

	"movie-api/internal/db"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func DeleteMovie(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	log.Println("RAW PATH:", req.Path)

	// Clean and split path
	path := strings.Trim(req.Path, "/")
	parts := strings.Split(path, "/")

	// Expected: ["movies", "{movieId}"]
	if len(parts) != 2 || parts[0] != "movies" {
		return events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       "movieId is required in path",
		}, nil
	}

	movieID := parts[1]

	_, err := db.Client.DeleteItem(context.Background(), &dynamodb.DeleteItemInput{
		TableName: aws.String("Movies"),
		Key: map[string]types.AttributeValue{
			"movieId": &types.AttributeValueMemberS{Value: movieID},
		},
	})
	if err != nil {
		log.Println("Delete error:", err)
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "Failed to delete movie",
		}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "Movie deleted successfully",
	}, nil
}

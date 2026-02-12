package handlers

import (
	"context"
	"encoding/json"
	"log"

	"movie-api/internal/db"
	"movie-api/internal/models"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func SearchMovies(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	resp, err := db.Client.Scan(context.Background(), &dynamodb.ScanInput{
		TableName: aws.String("Movies"),
	})
	if err != nil {
		log.Println("Scan error:", err)
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "Failed to scan movies",
		}, nil
	}

	movies := []models.Movie{}

	for _, item := range resp.Items {
		movie := models.Movie{
			MovieID:     item["movieId"].(*types.AttributeValueMemberS).Value,
			Title:       item["title"].(*types.AttributeValueMemberS).Value,
			Genre:       item["genre"].(*types.AttributeValueMemberS).Value,
			Year:        stringToInt(item["year"].(*types.AttributeValueMemberN).Value),
			Rating:      stringToFloat(item["rating"].(*types.AttributeValueMemberN).Value),
			Description: item["description"].(*types.AttributeValueMemberS).Value,
		}
		movies = append(movies, movie)
	}

	body, _ := json.Marshal(movies)

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(body),
	}, nil
}

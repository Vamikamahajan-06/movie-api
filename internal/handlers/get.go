package handlers

import (
	"context"
	"encoding/json"
	"log"
	"strings"

	"movie-api/internal/db"
	"movie-api/internal/models"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func GetMovie(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	log.Println("RAW PATH:", req.Path)

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

	resp, err := db.Client.GetItem(context.Background(), &dynamodb.GetItemInput{
		TableName: aws.String("Movies"),
		Key: map[string]types.AttributeValue{
			"movieId": &types.AttributeValueMemberS{Value: movieID},
		},
	})
	if err != nil {
		log.Println("GetItem error:", err)
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "Failed to get movie",
		}, nil
	}

	if resp.Item == nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 404,
			Body:       "Movie not found",
		}, nil
	}

	// Map DynamoDB item → Movie struct
	movie := models.Movie{
		MovieID:     resp.Item["movieId"].(*types.AttributeValueMemberS).Value,
		Title:       resp.Item["title"].(*types.AttributeValueMemberS).Value,
		Genre:       resp.Item["genre"].(*types.AttributeValueMemberS).Value,
		Year:        stringToInt(resp.Item["year"].(*types.AttributeValueMemberN).Value),
		Rating:      stringToFloat(resp.Item["rating"].(*types.AttributeValueMemberN).Value),
		Description: resp.Item["description"].(*types.AttributeValueMemberS).Value,
	}

	body, _ := json.Marshal(movie)

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(body),
	}, nil
}

package handlers

import (
	"context"
	"encoding/json"

	"movie-api/internal/db"
	"movie-api/internal/models"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/google/uuid"
)

func CreateMovie(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	// Parse JSON body
	var movie models.Movie
	err := json.Unmarshal([]byte(req.Body), &movie)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       "Invalid JSON body",
		}, nil
	}

	// Generate MovieID
	movie.MovieID = uuid.New().String()

	// Convert to DynamoDB item
	item := map[string]types.AttributeValue{
		"movieId":     &types.AttributeValueMemberS{Value: movie.MovieID},
		"title":       &types.AttributeValueMemberS{Value: movie.Title},
		"genre":       &types.AttributeValueMemberS{Value: movie.Genre},
		"year":        &types.AttributeValueMemberN{Value: intToString(movie.Year)},
		"rating":      &types.AttributeValueMemberN{Value: floatToString(movie.Rating)},
		"description": &types.AttributeValueMemberS{Value: movie.Description},
	}

	// Put into DynamoDB
	_, err = db.Client.PutItem(context.Background(), &dynamodb.PutItemInput{
		TableName: aws.String("Movies"),
		Item:      item,
	})
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "DynamoDB PutItem failed",
		}, nil
	}

	// Return created movie
	respBody, _ := json.Marshal(movie)

	return events.APIGatewayProxyResponse{
		StatusCode: 201,
		Body:       string(respBody),
	}, nil
}

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

func UpdateMovie(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	log.Println("RAW PATH:", req.Path)

	parts := strings.Split(strings.Trim(req.Path, "/"), "/")
	if len(parts) != 2 || parts[0] != "movies" {
		return events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       "movieId is required in path",
		}, nil
	}

	movieID := parts[1]

	var movie models.Movie
	err := json.Unmarshal([]byte(req.Body), &movie)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       "Invalid JSON body",
		}, nil
	}

	_, err = db.Client.UpdateItem(context.Background(), &dynamodb.UpdateItemInput{
		TableName: aws.String("Movies"),
		Key: map[string]types.AttributeValue{
			"movieId": &types.AttributeValueMemberS{Value: movieID},
		},
		UpdateExpression: aws.String("SET title = :title, genre = :genre, #yr = :year, rating = :rating, description = :description"),
		ExpressionAttributeNames: map[string]string{
			"#yr": "year",
		},
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":title":       &types.AttributeValueMemberS{Value: movie.Title},
			":genre":       &types.AttributeValueMemberS{Value: movie.Genre},
			":year":        &types.AttributeValueMemberN{Value: intToString(movie.Year)},
			":rating":      &types.AttributeValueMemberN{Value: floatToString(movie.Rating)},
			":description": &types.AttributeValueMemberS{Value: movie.Description},
		},
	})
	if err != nil {
		log.Println("UpdateItem error:", err)
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "Failed to update movie",
		}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "Movie updated successfully",
	}, nil
}

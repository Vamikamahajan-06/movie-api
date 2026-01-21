package handlers

import (
	"github.com/aws/aws-lambda-go/events"
)

func SearchMovies(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "SearchMovies called successfully",
	}, nil
}

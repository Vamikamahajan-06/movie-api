package handlers

import (
	"github.com/aws/aws-lambda-go/events"
)

func CreateMovie(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		StatusCode: 201,
		Body:       "CreateMovie called successfully",
	}, nil
}

package handlers

import (
	"github.com/aws/aws-lambda-go/events"
)

func UpdateMovie(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "UpdateMovie called successfully",
	}, nil
}

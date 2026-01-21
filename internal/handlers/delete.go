package handlers

import (
	"github.com/aws/aws-lambda-go/events"
)

func DeleteMovie(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "DeleteMovie called successfully",
	}, nil
}

package handlers

import (
	"github.com/aws/aws-lambda-go/events"
)

func GetMovie(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "GetMovie called successfully",
	}, nil
}

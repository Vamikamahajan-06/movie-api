package handlers

import (
	"context"

	"github.com/aws/aws/lambda-go/events"
)

func Router(ctx context.Context, req events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	switch req.RequestContext.HTTP.Method {
	case "POST":
		return CreateMovie(req)
	case "GET":
		if req.PathParameters["id"] != "" {
			return GetMovie(req)
		}
		return SearchMovies(req)

	case "PUT":
		return UpdateMovie(req)

	case "DELETE":
		return DeleteMovie(req)

	default:
		return events.APIGatewayV2HTTPResponse{
			StatusCode: 405,
			Body:       "Method not Allowed",
		}, nil
	}
}

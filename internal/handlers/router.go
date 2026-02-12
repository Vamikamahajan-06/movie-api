package handlers

import (
	"context"
	"strings"

	"github.com/aws/aws-lambda-go/events"
)

func Router(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	method := req.HTTPMethod
	path := req.Path

	switch {

	// Create movie
	case method == "POST" && path == "/movies":
		return CreateMovie(req)

	// Search movies (GET /movies?title=...&genre=...)
	case method == "GET" && path == "/movies":
		return SearchMovies(req)

	// Get movie by ID (GET /movies/{id})
	case method == "GET" && strings.HasPrefix(path, "/movies/"):
		return GetMovie(req)

	// Update movie (PUT /movies/{id})
	case method == "PUT" && strings.HasPrefix(path, "/movies/"):
		return UpdateMovie(req)

	// Delete movie (DELETE /movies/{id})
	case method == "DELETE" && strings.HasPrefix(path, "/movies/"):
		return DeleteMovie(req)

	default:
		return events.APIGatewayProxyResponse{
			StatusCode: 404,
			Body:       "Route not found",
		}, nil
	}
}

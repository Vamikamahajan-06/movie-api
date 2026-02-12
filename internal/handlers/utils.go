package handlers

import (
	"net/http"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
)

func intToString(i int) string {
	return strconv.Itoa(i)
}

func floatToString(f float64) string {
	return strconv.FormatFloat(f, 'f', -1, 64)
}

func serverError(err error) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusInternalServerError,
		Body:       err.Error(),
	}, nil
}

func awsString(s string) *string {
	return &s
}

func stringToInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func stringToFloat(s string) float64 {
	f, _ := strconv.ParseFloat(s, 64)
	return f
}

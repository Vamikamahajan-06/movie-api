package db

import (
	"context"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

var Client *dynamodb.Client

func Init() error {
	cfg, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		return err
	}
	Client = dynamodb.NewFromConfig(cfg)
	return nil
}
func TableName() string {
	return os.Getenv("Movies")

}

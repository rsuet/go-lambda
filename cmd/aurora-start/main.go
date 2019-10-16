package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws/session"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/rds"
)

type StartParams struct {
	DBClusterIdentifiers []string `json:"db_cluster_identifiers"`
}

func main() {
	lambda.Start(HandleRequest)
}

func HandleRequest(ctx context.Context, params StartParams) error {
	log.Print("Event Received!")
	service := rds.New(session.New())

	for _, identifier := range params.DBClusterIdentifiers {
		log.Print(fmt.Sprintf("Start Target: %v", identifier))
		input := &rds.StartDBClusterInput{
			DBClusterIdentifier: aws.String(identifier),
		}
		_, err := service.StartDBCluster(input)
		if err != nil {
			log.Print(err)
			return err
		}
	}

	log.Print("Event Finished!")
	return nil
}

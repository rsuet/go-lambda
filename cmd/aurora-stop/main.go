package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/service/rds"
)

type StopParams struct {
	DBClusterIdentifiers []string `json:"db_cluster_identifiers"`
}

func main() {
	lambda.Start(HandleRequest)
}

func HandleRequest(ctx context.Context, params StopParams) error {
	log.Print("Event Received!")
	service := rds.New(session.New())

	for _, identifier := range params.DBClusterIdentifiers {
		log.Print(fmt.Sprintf("Stop Target: %v", identifier))
		input := &rds.StopDBClusterInput{
			DBClusterIdentifier: aws.String(identifier),
		}
		_, err := service.StopDBCluster(input)
		if err != nil {
			log.Print(err)
			return err
		}
	}

	log.Print("Event Finished!")
	return nil
}

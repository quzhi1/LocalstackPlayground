package main

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

type Item struct {
	Key   string `json:"Key"`
	Value string `json:"value"`
}

func main() {
	// Openstack endpoint and region
	endpoint := "http://localhost:4566"
	region := "us-east-1"

	// Create Dynamodb AWS session
	config := &aws.Config{
		Endpoint: &endpoint,
		Region:   &region,
	}

	// Create a session
	var svc dynamodbiface.DynamoDBAPI
	sess := session.Must(session.NewSession(config))
	svc = dynamodb.New(sess)

	// Create item
	item := Item{
		Key:   "abc",
		Value: "Hello world",
	}
	putInput := &dynamodb.PutItemInput{
		Item: map[string]*dynamodb.AttributeValue{
			"Key": {
				S: &item.Key,
			}, "Value": {
				S: &item.Value,
			},
		},
		TableName: aws.String("ParameterAPI"),
	}
	_, err := svc.PutItem(putInput)
	if err != nil {
		panic(err)
	}

	// Read from dynamo
	result, err := svc.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String("ParameterAPI"),
		Key: map[string]*dynamodb.AttributeValue{
			"Key": {
				S: aws.String("abc"),
			},
		},
	})
	if err != nil {
		log.Fatalf("Got error calling GetItem: %s", err)
	}

	// Print result
	fmt.Println("Result from dynamo: key: " + *result.Item["Key"].S + ", value: " + *result.Item["Value"].S)
}

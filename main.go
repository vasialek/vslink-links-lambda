package main

import (
	"encoding/json"
	"log"
	"vslink-links-lambda/data"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(handler)
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Println("Loading all links...")

	repository := data.NewLinkRepository()
	links, err := repository.GetAllLinks()
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 500}, nil
	}

	json, _ := json.Marshal(links)

	response := events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers:    map[string]string{"Content-Type": "application/json"},
		Body:       string(json),
	}

	return response, nil
}

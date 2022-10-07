package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
)

func apiResponse(status int, body interface{}) (*events.APIGatewayProxyResponse, error) {
	stringBody, err := json.Marshal(body)
	if err != nil {
		fmt.Sprintf("an error occured: %v", err)
	}
	resp := events.APIGatewayProxyResponse{
		StatusCode: status,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: string(stringBody),
	}

	return &resp, err
}

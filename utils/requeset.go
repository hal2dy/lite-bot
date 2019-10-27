package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

// Response is of type APIGatewayProxyResponse since we're leveraging the
// AWS Lambda Proxy Request functionality (default behavior)
// https://serverless.com/framework/docs/providers/aws/events/apigateway/#lambda-proxy-integration
type Response events.APIGatewayProxyResponse

// BuildCommandData is data that need to trigger build
type BuildCommandData struct {
	Secret     string            `json:"secret"`
	Team       string            `json:"team"`
	Command    string            `json:"command"`
	UserName   string            `json:"username"`
	Parameters map[string]string `json:"parameters"`
}

// MakeResponse reponse the bot requeset
func MakeResponse(status int, response string) Response {
	var buf bytes.Buffer
	data, _ := json.Marshal(map[string]interface{}{"text": response})
	json.HTMLEscape(&buf, data)
	return Response{
		StatusCode:      200,
		IsBase64Encoded: false,
		Body:            buf.String(),
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}
}

// Post a request to url endpoint with provided body data
func Post(url string, data BuildCommandData) string {
	bodyJSON, _ := json.Marshal(data)
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(bodyJSON))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	_, err := client.Do(req)
	if err != nil {
		return fmt.Sprintf(`%s`, err)
	}
	return ""
}

package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// Response is of type APIGatewayProxyResponse since we're leveraging the
// AWS Lambda Proxy Request functionality (default behavior)
//
// https://serverless.com/framework/docs/providers/aws/events/apigateway/#lambda-proxy-integration
type Response events.APIGatewayProxyResponse

// var config struct {
// 	Secret  string
// 	WebHook string
// }

// SlackData is data that will be post from slack slash command
type SlackData struct {
	Token          string `json:"token"`
	TeamID         string `json:"team_id"`
	TeamDomain     string `json:"team_domain"`
	EnterpriseID   string `json:"enterprise_id"`
	EnterpriseName string `json:"enterprise_name"`
	ChannelID      string `json:"channel_id"`
	ChannelName    string `json:"channel_name"`
	UserID         string `json:"user_id"`
	UserName       string `json:"user_name"`
	Command        string `json:"command"`
	Text           string `json:"text"`
	ResponseURL    string `json:"response_url"`
	TriggerID      string `json:"trigger_id"`
}

// ParseSlackDataFromRequest is to parse request from Slack slash command to SlackData
func ParseSlackDataFromRequest(req string) SlackData {
	var slackData SlackData

	parseQuery, _ := url.ParseQuery(req)
	// convert map Token=[xxx] to Token=xxx
	m := make(map[string]string)
	for k, v := range parseQuery {
		m[k] = v[0]
	}

	jsonString, _ := json.Marshal(m)
	json.Unmarshal([]byte(jsonString), &slackData)
	return slackData
}

// MakeResponse is make the APIGatewayProxyResponse repsone
func MakeResponse(status int, response map[string]interface{}) Response {
	var buf bytes.Buffer
	data, _ := json.Marshal(response)
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

// ActionHelp will return result text of help command
func ActionHelp() map[string]interface{} {
	msg := make(map[string]interface{})
	msg["text"] = "*Lite bot Supported command*:\n" +
		"`help`: view this help `/lite help`\n" +
		"`build`: Build Lite to instance. `/lite build sg en 10000 primary`\n" +
		"> First 3 params is mandatory: *country, language, instance number without SHOP*\n" +
		"> Last params is optional : *primary|secondary* for multi-languages venture (default=primary)"

	return msg
}

// ActionBuild will return result text of build command
func ActionBuild(slackData SlackData, command []string) map[string]interface{} {
	if len(command) < 4 {
		msg := make(map[string]interface{})
		msg["text"] = "Missing mandatory params for Lite build, required: `country`, `language`, `instance number without SHOP`"
		return msg
	}

	secret := "secret"
	webHook := "https://zalora.io:8084/webhooks/webhook/lite"

	parameters := make(map[string]string)
	parameters["country"] = command[1]
	parameters["language"] = command[2]
	parameters["instance"] = command[3]
	if len(command) > 4 {
		parameters["option"] = command[4]
	} else {
		parameters["option"] = "primary"
	}
	if len(command) > 5 {
		parameters["branch"] = command[5]
	} else {
		parameters["branch"] = "master"
	}

	requestBody := make(map[string]interface{})
	requestBody["secret"] = secret
	requestBody["team"] = slackData.TeamDomain
	requestBody["command"] = slackData.Command
	requestBody["username"] = slackData.UserName
	requestBody["parameters"] = parameters

	requestJSONString, _ := json.Marshal(requestBody)
	req, err := http.NewRequest("POST", webHook, bytes.NewBuffer(requestJSONString))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	_, err = client.Do(req)
	if err != nil {
		panic(err)
	}
	// defer resp.Body.Close()
	msg := make(map[string]interface{})
	msg["text"] = fmt.Sprintf(
		`We got your request to build Lite on *SHOP-%s*, we will get back to you :wave:`,
		parameters["instance"])
	return msg
}

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (Response, error) {
	slackData := ParseSlackDataFromRequest(request.Body)
	command := strings.Fields(slackData.Text)
	if len(command) < 1 {
		msg := ActionHelp()
		return MakeResponse(400, msg), nil
	}

	var text map[string]interface{}
	switch action := command[0]; action {
	case "help":
		text = ActionHelp()
	case "build":
		text = ActionBuild(slackData, command)
	default:
		text = ActionHelp()
	}

	resp := MakeResponse(200, text)
	return resp, nil
}

// func init() {
// 	confPath := "config.json"
// 	file, _ := os.Open(confPath)
// 	decoder := json.NewDecoder(file)
// 	decoder.Decode(&config)
// }

func main() {
	lambda.Start(Handler)
}

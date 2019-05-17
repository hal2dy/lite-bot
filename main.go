package main

import (
	"context"
	"strings"

	"lite-bot/utils"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (utils.Response, error) {
	botData := utils.ParseBotData(request.Body)
	commands := strings.Fields(botData.Text)

	if len(commands) < 1 {
		return utils.MakeResponse(400, ActionHelp()), nil
	}

	result := ""
	action := commands[0]
	params := commands[1:]
	switch action {
	case "help":
		result = ActionHelp()
	case "build":
		result = ActionBuild(botData, params)
	default:
		result = ActionHelp()
	}

	resp := utils.MakeResponse(200, result)
	return resp, nil
}

func main() {
	lambda.Start(Handler)
}

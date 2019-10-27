package main

import (
	"context"
	"fmt"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/kelseyhightower/envconfig"

	"github.com/hal2dy/lite-bot/utils"
)

type Config struct {
	Webhook string
	Secret  string
}

// ActionHelp will return result text of help command
func ActionHelp() string {
	return "`/lite build [instance-number] [code-branch] [venture] [language-option]`\n" +
		"> *instance-number* (mandatory): SHOP instance number without SHOP- \n" +
		"> *code-branch* (optional): code branch that use to build, default is _master_\n" +
		"> *venture* (optional): auto choose current venture of the _instance_ if not provided\n" +
		"> *language-option* (optional): primary|secondary, default is _primary_\n" +
		"Example:\n" +
		"> `/lite build 15204` - Build Lite for instance _15204_\n" +
		"> `/lite build 15204 master sg secondary` -  Build Lite for instance _15204_ using _master_ branch code base for _sg_ venture as _secondary_ language"
}

// ActionBuild will return result text of build command
func ActionBuild(teamDomain string, username string, command string, commands []string) string {
	config := Config{}
	if err := envconfig.Process("LITE_SLACK_BOT", &config); err != nil {
		return "Config error - please let @litedevs know :sweat_smile:"
	}

	buildParams := utils.ParseBuildParams(commands)

	if buildParams["instance"] == "" {
		return "Invalid command, please check `/lite help`"
	}

	if buildParams["country"] == "" {
		buildParams["country"] = utils.GetInstanceVenture(buildParams["instance"])
	}
	buildParams["language"] = utils.GetLanguageFromVenture(buildParams["country"], buildParams["option"])

	buildCommand := utils.BuildCommandData{
		Secret:     config.Secret,
		Team:       teamDomain,
		Command:    command,
		UserName:   username,
		Parameters: buildParams,
	}

	if err := utils.Post(config.Webhook, buildCommand); err != "" {
		return fmt.Sprintf(`Error occur %s`, err)
	}

	return fmt.Sprintf(
		`Receive your request to build Lite on *SHOP-%s*, will get back to you :wave:`,
		buildParams["instance"])
}

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (utils.Response, error) {
	// a sample request.Body from Slack slash command
	// token=xxx&team_id=xxx&team_domain=hal2dy&channel_id=xxx&channel_name=privategroup&user_id=xxx&user_name=xxx&command=%2Flite&text=build+18801+master&response_url=xxx&trigger_id=xxx

	botData := utils.ParseSlackCommandData(request.Body)
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
		result = ActionBuild(botData.TeamDomain, botData.UserName, botData.Command, params)
	default:
		result = ActionHelp()
	}

	resp := utils.MakeResponse(200, result)
	return resp, nil
}

func main() {
	lambda.Start(Handler)
}

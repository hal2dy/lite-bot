package main

import (
	"fmt"
	"net/http"

	"lite-bot/utils"
)

// ActionHelp will return result text of help command
func ActionHelp() string {
	return "`/lite build [instance-number] [code-branch] [venture] [language-option]`\n" +
		"> *instance-number* (madatory): SHOP instance number without SHOP- \n" +
		"> *code-branch* (optional): code branch that use to build, default is _master_\n" +
		"> *venture* (optional): auto choose current venture of the _instance_ if not provided\n" +
		"> *language-option* (optional): primary|secondary, default is _primary_\n" +
		"Example:\n" +
		"> `/lite build 15204` - Build Lite for instance _15204_\n" +
		"> `/lite build 15204 master sg secondary` -  Build Lite for instance _15204_ using _master_ branch code base for _sg_ venture as _secondary_ language"
}

func connected() string {
	_, err := http.Get("https://www.google.com/")
	if err != nil {
		return "internet false"
	}
	return "internet true"
}

// ActionBuild will return result text of build command
func ActionBuild(botData utils.BotData, commands []string) string {
	secret := "xxx"
	webHook := "https://spinnaker.zalora.com/webhooks/"

	params := parseParams(commands)

	if params["instance"] == "" {
		return "Invalid command, please check `/lite help`"
	}

	if params["country"] == "" {
		params["country"] = utils.GetInstanceVenture(params["instance"])
	}
	params["language"] = utils.GetLanguageFromVenture(params["country"], params["option"])

	requestBody := make(map[string]interface{})
	requestBody["secret"] = secret
	requestBody["team"] = botData.TeamDomain
	requestBody["command"] = botData.Command
	requestBody["username"] = botData.UserName
	requestBody["parameters"] = params
	err := utils.Post(webHook, requestBody)
	if err != "" {
		return fmt.Sprintf(`Error occur %s`, err)
	}

	return fmt.Sprintf(
		`Receive your request to build Lite on *SHOP-%s*, will get back to you :wave:`,
		params["instance"])
}

func parseParams(commands []string) map[string]string {
	params := make(map[string]string)
	if len(commands) > 0 {
		params["instance"] = commands[0]
	} else {
		params["instance"] = ""
	}
	if len(commands) > 1 {
		params["branch"] = commands[1]
	} else {
		params["branch"] = "master"
	}
	if len(commands) > 2 {
		params["country"] = commands[2]
	} else {
		params["country"] = ""
	}
	if len(commands) > 3 {
		params["option"] = commands[3]
	} else {
		params["option"] = "primary"
	}
	return params
}

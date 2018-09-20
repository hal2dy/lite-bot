package main

import (
	"fmt"
	"net/http"

	"lite-bot/utils"
)

// ActionHelp will return result text of help command
func ActionHelp() string {
	return "_Lite bot Supported command_:\n" +
		"*help*: view this help `/lite help`\n" +
		"*build*: `/lite build instance-number [language-option] [code-branch] [venture]`\n" +
		"> - language-option has 2 option: primary|secondary, default is _primary_\n" +
		"> - code-branch is the branch in code base that use to build, default is _master_\n" +
		"> - venture if not provided will auto choose current venture of the _instance_\n" +
		"> `/lite build 15204` - Build Lite for instance _15204_\n" +
		"> `/lite build 15204 secondary master sg` -  Build Lite for instance _15204_ with _secondary_ language option using _master_ branch on Lite code base and build for _sg_ venture"
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
	secret := "some-secret"
	webHook := "https://zalora.io/webhooks/"

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
		`We got your request to build Lite on *SHOP-%s*, we will get back to you :wave:`,
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
		params["option"] = commands[1]
	} else {
		params["option"] = "primary"
	}
	if len(commands) > 2 {
		params["branch"] = commands[2]
	} else {
		params["branch"] = "master"
	}
	if len(commands) > 3 {
		params["country"] = commands[3]
	} else {
		params["country"] = ""
	}
	return params
}

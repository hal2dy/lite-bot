package utils

import (
	"encoding/json"
	"net/url"
)

// SlackCommandData is data that will be post from slack slash command
type SlackCommandData struct {
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

// ParseSlackCommandData is to parse request from Slack slash command
func ParseSlackCommandData(req string) (result SlackCommandData) {
	parseQuery, _ := url.ParseQuery(req)

	// convert map Token=[xxx] to Token=xxx
	normalizeParseQuery := make(map[string]string)
	for k, v := range parseQuery {
		normalizeParseQuery[k] = v[0]
	}

	jsonString, _ := json.Marshal(normalizeParseQuery)
	json.Unmarshal([]byte(jsonString), &result)

	return result
}

// ParseParams use to parse build parametes from request command
func ParseBuildParams(commands []string) map[string]string {
	params := map[string]string{
		"instance": "",
		"branch":   "master",
		"country":  "",
		"option":   "primary",
	}

	if len(commands) > 0 {
		params["instance"] = commands[0]
	}
	if len(commands) > 1 {
		params["branch"] = commands[1]
	}
	if len(commands) > 2 {
		params["country"] = commands[2]
	}
	if len(commands) > 3 {
		params["option"] = commands[3]
	}

	return params
}

package utils

import (
	"encoding/json"
	"net/url"
)

// BotData is data that will be post from slack slash command
type BotData struct {
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

// ParseBotData is to parse request from Slack slash command
func ParseBotData(req string) BotData {
	var BotData BotData

	parseQuery, _ := url.ParseQuery(req)
	// convert map Token=[xxx] to Token=xxx
	m := make(map[string]string)
	for k, v := range parseQuery {
		m[k] = v[0]
	}

	jsonString, _ := json.Marshal(m)
	json.Unmarshal([]byte(jsonString), &BotData)
	return BotData
}

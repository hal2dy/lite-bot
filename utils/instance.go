package utils

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

const instanceURLFormat = "http://instance-url.%s.com"

func isTitleElement(n *html.Node) bool {
	return n.Type == html.ElementNode && n.Data == "title"
}

func traverse(n *html.Node) (string, bool) {
	if isTitleElement(n) {
		return n.FirstChild.Data, true
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		result, ok := traverse(c)
		if ok {
			return result, ok
		}
	}
	return "", false
}

func getHTMLTitle(r io.Reader) (string, bool) {
	doc, err := html.Parse(r)
	if err != nil {
		return "", false
	}
	return traverse(doc)
}

func getVenture(title string) string {
	if strings.Contains(title, "Singapore") {
		return "sg"
	}
	if strings.Contains(title, "Malaysia") {
		return "my"
	}
	if strings.Contains(title, "Indonesia") {
		return "id"
	}
	if strings.Contains(title, "Philippines") {
		return "ph"
	}
	if strings.Contains(title, "臺灣") {
		return "tw"
	}
	if strings.Contains(title, "香港") || strings.Contains(title, "Hong Kong") {
		return "hk"
	}
	return "sg"
}

// GetLanguageFromVenture return language of the given venture
func GetLanguageFromVenture(venture string, languageOption string) string {
	switch venture {
	case "sg":
	case "my":
	case "ph":
		return "en"
	case "id":
		return "id"
	case "tw":
		return "zh"
	case "hk":
		if languageOption == "secondary" {
			return "zh"
		}
		return "en"
	}
	return "en"
}

// GetInstanceVenture return venture of the given instance
func GetInstanceVenture(instanceNumber string) string {
	venture := ""
	resp, err := http.Get(fmt.Sprintf(instanceURLFormat, instanceNumber))

	if err == nil {
		defer resp.Body.Close()
		if title, ok := getHTMLTitle(resp.Body); ok {
			venture = getVenture(title)
		}
	}

	return venture
}

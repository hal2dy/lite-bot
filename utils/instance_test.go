package utils

import "testing"

func TestGetVentureTW(t *testing.T) {
	result := getVenture("臺灣")
	if result != "tw" {
		t.Error("Expected venture is tw, got ", result)
	}
}

func TestGetVentureHK(t *testing.T) {
	result := getVenture("Hong Kong")
	if result != "hk" {
		t.Error("Expected venture is hk, got ", result)
	}
}

func TestGetVentureDefault(t *testing.T) {
	result := getVenture("")
	if result != "sg" {
		t.Error("Expected venture is sg, got ", result)
	}
}

func TestGetLanguageMY(t *testing.T) {
	result := GetLanguageFromVenture("my", "")
	if result != "en" {
		t.Error("Expected Language is en, got ", result)
	}
}
func TestGetLanguageID(t *testing.T) {
	result := GetLanguageFromVenture("id", "")
	if result != "id" {
		t.Error("Expected Language is id, got ", result)
	}
}

func TestGetLanguageHK(t *testing.T) {
	result := GetLanguageFromVenture("hk", "secondary")
	if result != "zh" {
		t.Error("Expected Language is zh, got ", result)
	}
}

func TestGetLanguageDefault(t *testing.T) {
	result := GetLanguageFromVenture("", "")
	if result != "en" {
		t.Error("Expected Language is en, got ", result)
	}
}

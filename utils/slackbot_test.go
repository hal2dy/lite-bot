package utils

import (
	"testing"
)

func TestParseParamsWithoutData(t *testing.T) {
	commands := []string{}
	result := ParseBuildParams(commands)
	if result["instance"] != "" {
		t.Error("Expected instance is empty, got ", result["instance"])
	}
	if result["option"] != "primary" {
		t.Error("Expected option is empty, got ", result["option"])
	}
	if result["branch"] != "master" {
		t.Error("Expected branch is master, got ", result["branch"])
	}
	if result["country"] != "" {
		t.Error("Expected country is empty, got ", result["country"])
	}
}

func TestParseParamsWithOnlyInstanceData(t *testing.T) {
	commands := []string{"12345"}
	result := ParseBuildParams(commands)
	if result["instance"] != "12345" {
		t.Error("Expected instance is 12345, got ", result["instance"])
	}
	if result["option"] != "primary" {
		t.Error("Expected option is empty, got ", result["option"])
	}
	if result["branch"] != "master" {
		t.Error("Expected branch is master, got ", result["branch"])
	}
	if result["country"] != "" {
		t.Error("Expected country is empty, got ", result["country"])
	}
}

func TestParseParamsWithOptionAndBranchData(t *testing.T) {
	commands := []string{"12345", "test-branch", "sg"}
	result := ParseBuildParams(commands)
	if result["instance"] != "12345" {
		t.Error("Expected instance is 12345, got ", result["instance"])
	}
	if result["option"] != "primary" {
		t.Error("Expected option is primary, got ", result["option"])
	}
	if result["branch"] != "test-branch" {
		t.Error("Expected branch is test-branch, got ", result["branch"])
	}
	if result["country"] != "sg" {
		t.Error("Expected country is sg, got ", result["country"])
	}
}

func TestParseParamsWithAllData(t *testing.T) {
	commands := []string{"12345", "test-branch", "vn", "primary"}
	result := ParseBuildParams(commands)
	if result["instance"] != "12345" {
		t.Error("Expected instance is 12345, got ", result["instance"])
	}
	if result["option"] != "primary" {
		t.Error("Expected option is primary, got ", result["option"])
	}
	if result["branch"] != "test-branch" {
		t.Error("Expected branch is test-branch, got ", result["branch"])
	}
	if result["country"] != "vn" {
		t.Error("Expected country is vn, got ", result["country"])
	}
}

package main

import "testing"

func TestParseParamsWithoutData(t *testing.T) {
	commands := []string{}
	result := parseParams(commands)
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
	result := parseParams(commands)
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
	commands := []string{"12345", "secondary", "test-branch"}
	result := parseParams(commands)
	if result["instance"] != "12345" {
		t.Error("Expected instance is 12345, got ", result["instance"])
	}
	if result["option"] != "secondary" {
		t.Error("Expected option is secondary, got ", result["option"])
	}
	if result["branch"] != "test-branch" {
		t.Error("Expected branch is test-branch, got ", result["branch"])
	}
	if result["country"] != "" {
		t.Error("Expected country is empty, got ", result["country"])
	}
}

func TestParseParamsWithAllData(t *testing.T) {
	commands := []string{"12345", "primary", "test-branch", "vn"}
	result := parseParams(commands)
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

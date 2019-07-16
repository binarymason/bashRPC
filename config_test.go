package main

import (
	"testing"

	. "github.com/binarymason/bashRPC/internal/testhelpers"
)

func TestLoad(t *testing.T) {
	cfg, _ := loadConfig("./test/data/simple_config.yml")
	expectedRoutes := []route{
		route{Path: "/foo", Cmd: "echo foo"},
		route{Path: "/bar", Cmd: "echo bar"},
		route{Path: "/date", Cmd: "date"},
	}
	Assert(cfg.Routes, expectedRoutes, t)
}

func TestValidConfig(t *testing.T) {
	When("config is valid")
	cfg := config{Port: "1234", Secret: "secret", Whitelisted: []string{"127.0.0.1"}}

	Then("there should be NO errors")
	if err := validateConfig(cfg); err != nil {
		t.Errorf("expected NO errors but received %v", err)
	}
}

func TestConfigMissingPort(t *testing.T) {
	When("port is not specified")
	cfg := config{Secret: "secret", Whitelisted: []string{"127.0.0.1"}}

	Then("an error is returned")
	if err := validateConfig(cfg); err == nil {
		t.Error("expected errors but received none")
	}
}

func TestConfigMissingSecret(t *testing.T) {
	When("secret is not specified")
	cfg := config{Port: "8675", Whitelisted: []string{"127.0.0.1"}}

	Then("an error is returned")
	if err := validateConfig(cfg); err == nil {
		t.Error("expected errors but received none")
	}
}

func TestConfigMissingWhitelisted(t *testing.T) {
	When("whitelisted clients are not specified")
	cfg := config{Secret: "secret", Port: "8675", Whitelisted: []string{}}

	Then("an error is returned")
	if err := validateConfig(cfg); err == nil {
		t.Error("expected errors but received none")
	}
}

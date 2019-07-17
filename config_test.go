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

func TestSetConfigDefaults(t *testing.T) {
	cfg := config{}

	setConfigDefaults(&cfg)

	Assert(cfg.Port, "8675", t)
	Assert(cfg.Key, "/etc/bashrpc/pki/bashrpc.key", t)
	Assert(cfg.Cert, "/etc/bashrpc/pki/bashrpc.cert", t)
}

var validConfig = config{
	Port:        "1234",
	Secret:      "secret",
	Whitelisted: []string{"127.0.0.1"},
	Key:         "/path/to/key",
	Cert:        "/path/to/cert",
}

func TestValidConfig(t *testing.T) {
	When("config is valid")
	Then("there should be NO errors")
	if err := validateConfig(validConfig); err != nil {
		t.Errorf("expected NO errors but received %v", err)
	}
}

func TestConfigMissingPort(t *testing.T) {
	When("port is not specified")
	cfg := validConfig
	cfg.Port = ""

	Then("an error is returned")
	if err := validateConfig(cfg); err == nil {
		t.Error("expected errors but received none")
	}
}

func TestConfigMissingSecret(t *testing.T) {
	When("secret is not specified")
	cfg := validConfig
	cfg.Secret = ""

	Then("an error is returned")
	if err := validateConfig(cfg); err == nil {
		t.Error("expected errors but received none")
	}
}

func TestConfigMissingWhitelisted(t *testing.T) {
	When("whitelisted clients are not specified")
	cfg := validConfig
	cfg.Whitelisted = []string{}

	Then("an error is returned")
	if err := validateConfig(cfg); err == nil {
		t.Error("expected errors but received none")
	}
}

func TestConfigMissingKey(t *testing.T) {
	When("key file is not specified")
	cfg := validConfig
	cfg.Key = ""

	Then("an error is returned")
	if err := validateConfig(cfg); err == nil {
		t.Error("expected errors but received none")
	}
}

func TestConfigMissingCert(t *testing.T) {
	When("cert file is not specified")
	cfg := validConfig
	cfg.Cert = ""

	Then("an error is returned")
	if err := validateConfig(cfg); err == nil {
		t.Error("expected errors but received none")
	}
}

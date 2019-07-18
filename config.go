package main

import (
	"errors"
	"io/ioutil"
	"strings"

	"gopkg.in/yaml.v2"
)

const defaultCertPath = "/etc/bashrpc/"

type config struct {
	Cert        string   `yaml:"cert"`
	Key         string   `yaml:"key"`
	Log         string   `yaml:"log"`
	Port        string   `yaml:"port"`
	Routes      []route  `yaml:"routes"`
	Secret      string   `yaml:"secret"`
	Whitelisted []string `yaml:"whitelisted_clients"`
}

type route struct {
	Path string `yaml:"path"`
	Cmd  string `yaml:"cmd"`
}

func loadConfig(p string) (config, error) {
	cfg := config{}
	setConfigDefaults(&cfg)

	data, err := ioutil.ReadFile(p)
	if err != nil {
		return cfg, err
	}

	err = yaml.Unmarshal([]byte(data), &cfg)
	return cfg, err
}

func setConfigDefaults(cfg *config) {
	defaultPKIPath := "/etc/bashrpc/pki"
	cfg.Cert = defaultPKIPath + "/bashrpc.cert"
	cfg.Key = defaultPKIPath + "/bashrpc.key"
	cfg.Log = "bashrpc.log"
	cfg.Port = "8675"
}

func validateConfig(cfg *config) (err error) {
	var issues []string

	if cfg.Port == "" {
		issues = append(issues, "port is missing")
	}

	if cfg.Secret == "" {
		issues = append(issues, "secret is missing")
	}

	if cfg.Key == "" {
		issues = append(issues, "key is missing")
	}

	if cfg.Cert == "" {
		issues = append(issues, "cert is missing")
	}

	if cfg.Log == "" {
		issues = append(issues, "log is missing")
	}

	if len(cfg.Whitelisted) == 0 {
		issues = append(issues, "no whitelisted clients are specified")
	}

	if len(issues) > 0 {
		err = errors.New("config validation errors: " + strings.Join(issues, ", "))
	}

	return
}

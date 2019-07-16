package main

import (
	"errors"
	"io/ioutil"
	"strings"

	"gopkg.in/yaml.v2"
)

type config struct {
	Port        string   `yaml:"port"`
	Secret      string   `yaml:"secret"`
	Whitelisted []string `yaml:"whitelisted_clients"`
	Routes      []route  `yaml:"routes"`
}

type route struct {
	Path string `yaml:"path"`
	Cmd  string `yaml:"cmd"`
}

func loadConfig(p string) (config, error) {
	cfg := config{}
	data, err := ioutil.ReadFile(p)
	if err != nil {
		return cfg, err
	}

	err = yaml.Unmarshal([]byte(data), &cfg)
	return cfg, err
}

func validateConfig(cfg config) error {
	var issues []string

	if cfg.Port == "" {
		issues = append(issues, "port is missing")
	}

	if cfg.Secret == "" {
		issues = append(issues, "secret is missing")
	}

	if len(cfg.Whitelisted) == 0 {
		issues = append(issues, "no whitelisted clients are specified")
	}

	if len(issues) > 0 {
		return errors.New("config validation errors: " + strings.Join(issues, ", "))
	}

	return nil
}

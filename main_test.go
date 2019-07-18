package main

import (
	"testing"

	. "github.com/binarymason/bashRPC/internal/testhelpers"
)

func TestCLIOverrides(t *testing.T) {
	cfg := config{
		Log:  "/path/to/log",
		Port: "8443",
	}

	opts := options{
		logPath: "/path/to/another/log",
		port:    "3210",
	}

	overrideConfigWithOptions(&cfg, opts)

	Assert(cfg.Log, "/path/to/another/log", t)
	Assert(cfg.Port, "3210", t)
}

package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

type options struct {
	configPath string
	logPath    string
	port       string
}

var opts options

func init() {
	c := flag.String("c", "", "specify bashRPC config file")
	p := flag.String("p", "", "specify bashRPC port")
	l := flag.String("log", "bashrpc.log", "specify log file")
	flag.Parse()

	opts = options{
		configPath: *c,
		logPath:    *l,
		port:       *p,
	}
}

func main() {
	if opts.configPath == "" {
		fmt.Println("config file argument is required")
		flag.Usage()
		os.Exit(1)
	}

	var (
		logFile *os.File
		cfg     config
		err     error
		rtr     router
	)

	if cfg, err = loadConfig(opts.configPath); err != nil {
		panic(err)
	}

	overrideConfigWithOptions(&cfg, opts)

	if err = validateConfig(&cfg); err != nil {
		panic(err)
	}

	if logFile, err = initLog(opts.logPath); err != nil {
		panic(err)
	}

	defer logFile.Close()
	log.SetOutput(io.MultiWriter(os.Stdout, logFile))
	log.Println("logging to", opts.logPath)

	if rtr, err = newRouter(cfg); err != nil {
		panic(err)
	}

	if err = rtr.listen(); err != nil {
		panic(err)
	}
}

func overrideConfigWithOptions(cfg *config, opts options) {
	if opts.logPath != "" {
		cfg.Log = opts.logPath
	}
	if opts.port != "" {
		cfg.Port = opts.port
	}
}

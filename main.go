package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

var (
	configPath string
	logPath    string
)

func init() {
	c := flag.String("c", "", "specify bashRPC config file")
	l := flag.String("log", "bashrpc.log", "specify log file")
	flag.Parse()

	configPath = *c
	logPath = *l
}

func main() {
	if configPath == "" {
		fmt.Println("config file argument is required")
		flag.Usage()
		os.Exit(1)
	}

	logFile, err := os.OpenFile(logPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		panic(err)
	}

	defer logFile.Close()
	log.SetOutput(io.MultiWriter(os.Stdout, logFile))

	rtr, err := newRouter(configPath)
	if err != nil {
		log.Fatal(fmt.Sprintf("%v", err))
		os.Exit(1)
	}

	if err := rtr.listen(); err != nil {
		log.Fatal(fmt.Sprintf("%v", err))
		os.Exit(1)
	}
}

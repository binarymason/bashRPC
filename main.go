package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

var configPath string

func init() {
	c := flag.String("c", "", "specify bashRPC config file")
	flag.Parse()
	configPath = *c
}

func main() {
	if configPath == "" {
		fmt.Println("config file argument is required")
		flag.Usage()
		os.Exit(1)
	}

	rtr, err := newRouter(configPath)
	if err != nil {
		log.Fatal(fmt.Sprintf("%v", err))
		os.Exit(1)
	}

	rtr.listen()
}

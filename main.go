package main

import (
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
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

type router struct {
	config config
}

var configPath string

func init() {
	c := flag.String("c", "", "go-remote config file")
	flag.Parse()
	configPath = *c
}

// TODO:
// * check authorized host
// * check authorized auth
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

func newRouter(p string) (router, error) {
	rtr := router{}
	cfg, err := loadConfig(p)

	if err != nil {
		return rtr, err
	}

	if err := validateConfig(cfg); err != nil {
		return rtr, err
	}

	rtr.config = cfg

	return rtr, nil
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

func (rtr *router) listen() {
	http.HandleFunc("/", rtr.handler)

	fmt.Println("listening on port", rtr.config.Port)
	http.ListenAndServe(":"+rtr.config.Port, nil)

}

func (rtr *router) handler(w http.ResponseWriter, r *http.Request) {
	if !rtr.authorizedRequest(r) {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	path := r.URL.Path

	route, err := rtr.routeForPath(path)

	if err != nil {
		http.Error(w, "invalid route", http.StatusNotFound)
		return
	}

	command, args := parseCommand(route.Cmd)

	out, err := exec.Command(command, args...).Output()
	if err != nil {
		http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
		return
	}

	w.Write(out)
}

func (rtr *router) authorizedRequest(r *http.Request) bool {
	ip := strings.Split(r.RemoteAddr, ":")[0] // remove port
	auth := r.Header.Get("Authorization")

	return validIP(ip, rtr.config.Whitelisted) && (auth == rtr.config.Secret)
}

func validIP(ip string, whitelisted []string) bool {
	for _, w := range whitelisted {
		if ip == w {
			return true
		}
	}

	return false
}

func (rtr *router) routeForPath(path string) (r route, err error) {
	for _, route := range rtr.config.Routes {
		if route.Path == path {
			return route, nil
		}
	}

	return r, errors.New("Route not found: " + path)
}

func parseCommand(s string) (c string, args []string) {
	command := strings.Split(s, " ")
	return command[0], command[1:]
}

package main

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

type router struct {
	config config
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

	out, err := runCommand(route.Cmd) // TODO: stream output
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

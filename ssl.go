package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// initSSL creates a SSL certificate and key using system's openssl.
func initSSL(certPath, keyPath string) ([]byte, error) {
	if out, err := initRndFile(); err != nil {
		return out, err
	}

	if out, err := initSSLKey(keyPath); err != nil {
		return out, err
	}

	return initSSLCert(certPath, keyPath)
}

func initSSLKey(keyPath string) (out []byte, err error) {
	if fileExists(keyPath) {
		log.Println("using existing SSL key:", keyPath)
		return
	}

	if err = mkdirP(keyPath); err != nil {
		return
	}

	log.Println("SSL key does not exist")
	log.Println("creating", keyPath)
	return runCommand(fmt.Sprintf("openssl genrsa -out %s 4096", keyPath))
}

func initSSLCert(certPath, keyPath string) (out []byte, err error) {
	if fileExists(certPath) {
		log.Println("using existing SSL cert:", certPath)
		return
	}

	var fqdn string
	fqdn, err = getFQDN()

	if err != nil {
		return
	}

	if err = mkdirP(certPath); err != nil {
		return
	}

	args := []string{
		"req",
		"-new",
		"-days",
		"3650",
		"-nodes",
		"-x509",
		"-key",
		keyPath,
		"-subj",
		fmt.Sprintf("/C=US/ST=Somewhere/L=Unknown/O=Idk/CN=%s", fqdn),
		"-out",
		certPath,
	}

	log.Println("SSL cert does not exist")
	log.Println("creating", certPath)
	return runCommand("openssl " + strings.Join(args, " "))
}

func initRndFile() (out []byte, err error) {
	rndPath := fmt.Sprintf("%s/.rnd", os.Getenv("HOME"))
	if fileExists(rndPath) {
		return
	}

	log.Println("openssl random seed file does not exist")
	log.Println("creating", rndPath)
	command := fmt.Sprintf("openssl rand -out %s -hex 256", rndPath)
	return runCommand(command)
}

func getFQDN() (fqdn string, err error) {
	out, err := runCommand("hostname --fqdn")

	if err != nil {
		return
	}

	fqdn = string(out)
	fqdn = fqdn[:len(fqdn)-1] // removing EOL

	return
}

func mkdirP(p string) error {
	absPath, _ := filepath.Abs(p)
	dir := filepath.Dir(absPath)
	_, err := os.Stat(absPath)

	if os.IsExist(err) {
		return nil
	}

	return os.MkdirAll(dir, 0700)
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false

	}
	return !info.IsDir()

}

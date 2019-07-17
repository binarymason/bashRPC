package main

import (
	"fmt"
	"strings"
)

func initSSL(outDir string) ([]byte, error) {
	if out, err := initRndFile(); err != nil {
		return out, err
	}

	fqdn, err := getFQDN()

	if err != nil {
		return []byte{}, err
	}

	keyPath := fmt.Sprintf("%s/%s.key", outDir, fqdn)
	certPath := fmt.Sprintf("%s/%s.cert", outDir, fqdn)
	command := "openssl"
	args := []string{
		"req",
		"-new",
		"-newkey",
		"rsa:4096",
		"-days",
		"3650",
		"-nodes",
		"-x509",
		"-subj",
		fmt.Sprintf("/C=US/ST=Somewhere/L=Unknown/O=Idk/CN=%s", fqdn),
		"-keyout",
		keyPath,
		"-out",
		certPath,
	}

	return runCommand(command + " " + strings.Join(args, " "))
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

func initRndFile() ([]byte, error) {
	return runCommand(`openssl rand -out "$HOME/.rnd" -hex 256`)
}

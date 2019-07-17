package main

import (
	"os"
	"os/exec"
	"testing"

	. "github.com/binarymason/bashRPC/internal/testhelpers"
	"github.com/pkg/errors"
)

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false

	}
	return !info.IsDir()

}
func TestInitSSL(t *testing.T) {
	var (
		keyPath  = "/tmp/test/test-host.key"
		certPath = "/tmp/test/test-host.cert"
	)
	Given("openssl is available on the machine")
	if out, err := exec.Command("openssl", "version").CombinedOutput(); err != nil {
		t.Error(errors.Wrap(err, string(out)))
	}

	When("an output directory is specified")
	if _, err := initSSL(certPath, keyPath); err != nil {
		t.Error(err)
	}

	Then("a SSL private key is generated")
	Assert(fileExists(keyPath), true, t)
	if out, err := exec.Command("openssl", "rsa", "-in", keyPath, "-check").CombinedOutput(); err != nil {
		t.Error(errors.Wrap(err, string(out)))
	}

	And("a SSL certificate is generated")
	Assert(fileExists(certPath), true, t)
	if out, err := exec.Command("openssl", "x509", "-in", certPath, "-text").CombinedOutput(); err != nil {
		t.Error(errors.Wrap(err, string(out)))
	}
}

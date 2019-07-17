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
	Given("openssl is available on the machine")
	if out, err := exec.Command("openssl", "version").CombinedOutput(); err != nil {
		t.Error(errors.Wrap(err, string(out)))
	}

	When("an output directory is specified")
	outDir := "/tmp/test"
	if _, err := initSSL(outDir); err != nil {
		t.Error(err)
	}

	Then("a SSL private key is generated")
	key := outDir + "/tp.key"
	Assert(fileExists(key), true, t)
	if out, err := exec.Command("openssl", "rsa", "-in", key, "-check").CombinedOutput(); err != nil {
		t.Error(errors.Wrap(err, string(out)))
	}

	And("a SSL certificate is generated")
	cert := outDir + "/tp.cert"
	Assert(fileExists(cert), true, t)
	if out, err := exec.Command("openssl", "x509", "-in", cert, "-text").CombinedOutput(); err != nil {
		t.Error(errors.Wrap(err, string(out)))
	}
}

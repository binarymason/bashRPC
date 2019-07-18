package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"testing"

	. "github.com/binarymason/bashRPC/internal/testhelpers"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

func createFile(path, content string) error {
	return ioutil.WriteFile(path, []byte(content), 0644)
}

func TestInitSSL(t *testing.T) {
	id, _ := uuid.NewUUID()

	var (
		testDir  = "/tmp/bashrpc-testing"
		pkiDir   = fmt.Sprintf("%s/test-%v", testDir, id)
		keyPath  = pkiDir + "/pki/test-host.key"
		certPath = pkiDir + "/pki/test-host.cert"
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

	os.RemoveAll("/tmp/bashrpc")
}

func TestExistingSSLKey(t *testing.T) {
	keyPath := "/tmp/test-ssl.key"
	Given("an SSL key that already exists")
	if err := createFile(keyPath, "I am a key"); err != nil {
		t.Error(err)
	}

	When("initializing SSL key")
	if _, err := initSSLKey(keyPath); err != nil {
		t.Error(err)
	}

	Then("it should NOT be overwritten")
	key, err := ioutil.ReadFile(keyPath)
	if err != nil {
		t.Error(err)
	}

	Assert(string(key), "I am a key", t)

	os.Remove(keyPath)
}

func TestExistingSSLCert(t *testing.T) {
	keyPath := "/tmp/test-ssl.key"
	certPath := "/tmp/test-ssl.cert"
	Given("an SSL cert that already exists")
	if err := createFile(keyPath, "I am a key"); err != nil {
		t.Error(err)
	}
	if err := createFile(certPath, "I am a cert"); err != nil {
		t.Error(err)
	}

	When("initializing SSL cert")
	if _, err := initSSLCert(certPath, keyPath); err != nil {
		t.Error(err)
	}

	Then("it should NOT be overwritten")
	cert, err := ioutil.ReadFile(certPath)
	if err != nil {
		t.Error(err)
	}

	Assert(string(cert), "I am a cert", t)

	os.Remove(keyPath)
	os.Remove(certPath)
}

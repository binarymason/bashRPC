package main

import (
	"testing"

	. "github.com/binarymason/bashRPC/internal/testhelpers"
	"github.com/pkg/errors"
)

func TestMultilineCommand(t *testing.T) {
	command := `
echo foo
echo bar
echo baz
  `

	out, err := runCommand(command)
	Assert(err, nil, t)
	Assert(string(out), "foo\nbar\nbaz\n", t)
}

func TestFailingCommand(t *testing.T) {
	out, err := runCommand("echo 'BOOM!' && exit 1")

	if err == nil {
		t.Error("expected an error but received none")
	}

	Assert(string(out), "BOOM!\n", t)
	expectedErr := errors.Wrap(errors.New("exit status 1"), "BOOM!\n")
	Assert(err, expectedErr, t)
}

func TestPipedCommand(t *testing.T) {
	command := `echo "it works with pipe" | grep pipe | awk '{ print $1 " " $2 }'`
	out, err := runCommand(command)
	Assert(err, nil, t)
	Assert(string(out), "it works\n", t)

}

package main

import (
	"os/exec"

	"github.com/pkg/errors"
)

func runCommand(command string) (out []byte, err error) {
	out, err = exec.Command("bash", "-ce", command).CombinedOutput()

	if err != nil {
		return out, errors.Wrap(err, string(out))
	}

	return
}

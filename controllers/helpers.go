package controllers

import (
	"bytes"
	"os/exec"
)

func CmdExec(args ...string) (*bytes.Buffer, error) {
	baseCmd := args[0]
	cmdArgs := args[1:]

	cmd := exec.Command(baseCmd, cmdArgs...)

	cmdOutput := &bytes.Buffer{}
	cmd.Stdout = cmdOutput
	cmdErr := &bytes.Buffer{}
	cmd.Stderr = cmdErr
	err := cmd.Run()
	if err != nil {
		return cmdErr, err
	}
	return cmdOutput, nil
}

func CmdExecStrOutput(args ...string) (string, error) {
	res, err := CmdExec(args...)
	return res.String(), err
}

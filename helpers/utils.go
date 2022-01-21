package helpers

import (
	"bytes"
	"os"
	"os/exec"
	"runtime"
)

func DirectoryExists(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return info.IsDir()
}

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

func SliceContains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func PrintPdf(path string, printer string) (string, error) {
	var err error
	var out string
	if runtime.GOOS == "windows" {
		out, err = CmdExecStrOutput("PDFtoPrinter", path, printer)
	} else {
		out, err = CmdExecStrOutput("lp", "-d", printer, path)
	}
	return out, err
}

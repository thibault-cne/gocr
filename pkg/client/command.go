package client

import (
	"bufio"
	"os/exec"
)

func startCommand(cmd *exec.Cmd) *bufio.Scanner {
	so, err := cmd.StdoutPipe()
	if err != nil {
		panic(err)
	}

	cmd.Start()

	stdout := bufio.NewScanner(so)

	return stdout
}

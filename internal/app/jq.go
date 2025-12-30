package app

import (
	"bytes"
	"errors"
	"os"
	"os/exec"
	"strings"
)

func runJQ(filter string, body []byte) error {
	filter = strings.TrimSpace(filter)
	if filter == "" {
		_, err := os.Stdout.Write(body)
		return err
	}
	if _, err := exec.LookPath("jq"); err != nil {
		return errors.New("jq not found")
	}
	cmd := exec.Command("jq", filter)
	cmd.Stdin = bytes.NewReader(body)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

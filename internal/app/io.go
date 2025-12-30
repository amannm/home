package app

import (
	"errors"
	"io"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

func readStdin() ([]byte, error) {
	return io.ReadAll(os.Stdin)
}

func readJSONFromFlags(cmd *cobra.Command, fileFlag, stdinFlag string) ([]byte, error) {
	useStdin := false
	if stdinFlag != "" {
		v, err := cmd.Flags().GetBool(stdinFlag)
		if err != nil {
			return nil, err
		}
		useStdin = v
	}
	if useStdin {
		return readStdin()
	}
	if fileFlag == "" {
		return nil, errors.New("no input")
	}
	path, err := cmd.Flags().GetString(fileFlag)
	if err != nil {
		return nil, err
	}
	if strings.TrimSpace(path) == "" {
		return nil, errors.New("file is required or use --stdin")
	}
	return os.ReadFile(path)
}

func splitQuery(pair string) (string, string, bool) {
	if strings.Contains(pair, "=") {
		parts := strings.SplitN(pair, "=", 2)
		return strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1]), parts[0] != ""
	}
	return "", "", false
}

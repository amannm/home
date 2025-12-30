package app

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"github.com/spf13/cobra"
)

func (a *App) Dist(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("dist: missing subcommand")
	}
	switch args[0] {
	case "info":
		return a.get(a.api("dist/getDistributionInfo"), nil)
	case "server":
		body, err := readJSONFromFlags(cmd, "file", "stdin")
		if err != nil {
			return err
		}
		return a.post(a.api("dist/setServerInfo"), nil, body)
	case "client":
		body, err := readJSONFromFlags(cmd, "file", "stdin")
		if err != nil {
			return err
		}
		return a.post(a.api("dist/setClientInfo"), nil, body)
	case "start":
		num, err := cmd.Flags().GetInt("num")
		if err != nil {
			return err
		}
		q := url.Values{}
		q.Set("num", fmt.Sprintf("%d", num))
		return a.get(a.api("dist/startDistribution"), q)
	case "stop":
		return a.get(a.api("dist/stopDistribution"), nil)
	case "group-name":
		useStdin, err := cmd.Flags().GetBool("stdin")
		if err != nil {
			return err
		}
		var body []byte
		if useStdin {
			body, err = readStdin()
			if err != nil {
				return err
			}
		} else {
			name, err := cmd.Flags().GetString("name")
			if err != nil {
				return err
			}
			if strings.TrimSpace(name) == "" {
				return fmt.Errorf("dist group-name: --name or --stdin is required")
			}
			body, err = json.Marshal(map[string]any{"name": name})
			if err != nil {
				return err
			}
		}
		return a.post(a.api("dist/setGroupName"), nil, body)
	default:
		return fmt.Errorf("dist: unknown command %s", args[0])
	}
}

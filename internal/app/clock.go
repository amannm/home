package app

import (
	"fmt"
	"net/url"

	"github.com/spf13/cobra"
)

func (a *App) Clock(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("clock: missing subcommand")
	}
	switch args[0] {
	case "settings":
		return a.get(a.api("clock/getSettings"), nil)
	case "auto-sync":
		enable, err := cmd.Flags().GetBool("enable")
		if err != nil {
			return err
		}
		q := url.Values{}
		q.Set("enable", fmt.Sprintf("%t", enable))
		return a.get(a.api("clock/setAutoSync"), q)
	case "datetime":
		dt, err := cmd.Flags().GetString("date-time")
		if err != nil {
			return err
		}
		q := url.Values{}
		q.Set("date_time", dt)
		return a.get(a.api("clock/setDateAndTime"), q)
	case "format":
		if len(args) < 2 {
			return fmt.Errorf("clock format: missing value")
		}
		q := url.Values{}
		q.Set("format", args[1])
		return a.get(a.api("clock/setClockFormat"), q)
	case "alarm":
		body, err := readJSONFromFlags(cmd, "file", "stdin")
		if err != nil {
			return err
		}
		return a.post(a.api("clock/setAlarmSettings"), nil, body)
	default:
		return fmt.Errorf("clock: unknown command %s", args[0])
	}
}

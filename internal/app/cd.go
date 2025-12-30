package app

import (
	"fmt"
	"net/url"
	"strconv"

	"github.com/spf13/cobra"
)

func (a *App) CD(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("cd: missing subcommand")
	}
	switch args[0] {
	case "play-info":
		return a.get(a.api("cd/getPlayInfo"), nil)
	case "playback":
		if len(args) < 2 {
			return fmt.Errorf("cd playback: missing value")
		}
		num, err := cmd.Flags().GetInt("num")
		if err != nil {
			return err
		}
		q := url.Values{}
		q.Set("playback", args[1])
		if cmd.Flags().Changed("num") {
			q.Set("num", strconv.Itoa(num))
		}
		return a.get(a.api("cd/setPlayback"), q)
	case "tray":
		return a.get(a.api("cd/toggleTray"), nil)
	case "repeat":
		if len(args) < 2 {
			return fmt.Errorf("cd repeat: missing value")
		}
		q := url.Values{}
		q.Set("mode", args[1])
		return a.get(a.api("cd/setRepeat"), q)
	case "shuffle":
		if len(args) < 2 {
			return fmt.Errorf("cd shuffle: missing value")
		}
		q := url.Values{}
		q.Set("mode", args[1])
		return a.get(a.api("cd/setShuffle"), q)
	case "repeat-toggle":
		return a.get(a.api("cd/toggleRepeat"), nil)
	case "shuffle-toggle":
		return a.get(a.api("cd/toggleShuffle"), nil)
	case "direct":
		enable, err := cmd.Flags().GetBool("enable")
		if err != nil {
			return err
		}
		q := url.Values{}
		q.Set("enable", strconv.FormatBool(enable))
		return a.get(a.api("cd/setDirect"), q)
	default:
		return fmt.Errorf("cd: unknown command %s", args[0])
	}
}

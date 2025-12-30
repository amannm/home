package app

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

func (a *App) Tuner(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("tuner: missing subcommand")
	}
	switch args[0] {
	case "preset-info":
		band, err := cmd.Flags().GetString("band")
		if err != nil {
			return err
		}
		q := url.Values{}
		q.Set("band", band)
		return a.get(a.api("tuner/getPresetInfo"), q)
	case "play-info":
		return a.get(a.api("tuner/getPlayInfo"), nil)
	case "band":
		if len(args) < 2 {
			return fmt.Errorf("tuner band: missing value")
		}
		q := url.Values{}
		q.Set("band", args[1])
		return a.get(a.api("tuner/setBand"), q)
	case "freq":
		band, err := cmd.Flags().GetString("band")
		if err != nil {
			return err
		}
		tuning, err := cmd.Flags().GetString("tuning")
		if err != nil {
			return err
		}
		num, err := cmd.Flags().GetInt("num")
		if err != nil {
			return err
		}
		q := url.Values{}
		q.Set("band", band)
		q.Set("tuning", tuning)
		if cmd.Flags().Changed("num") {
			q.Set("num", strconv.Itoa(num))
		}
		return a.get(a.api("tuner/setFreq"), q)
	case "recall":
		band, err := cmd.Flags().GetString("band")
		if err != nil {
			return err
		}
		num, err := cmd.Flags().GetInt("num")
		if err != nil {
			return err
		}
		q := url.Values{}
		q.Set("zone", zoneOrDefault(a.Options.Zone))
		q.Set("band", band)
		q.Set("num", strconv.Itoa(num))
		return a.get(a.api("tuner/recallPreset"), q)
	case "switch":
		dir, err := cmd.Flags().GetString("dir")
		if err != nil {
			return err
		}
		q := url.Values{}
		q.Set("dir", dir)
		return a.get(a.api("tuner/switchPreset"), q)
	case "store":
		num, err := cmd.Flags().GetInt("num")
		if err != nil {
			return err
		}
		q := url.Values{}
		q.Set("num", strconv.Itoa(num))
		return a.get(a.api("tuner/storePreset"), q)
	case "clear":
		band, err := cmd.Flags().GetString("band")
		if err != nil {
			return err
		}
		num, err := cmd.Flags().GetInt("num")
		if err != nil {
			return err
		}
		q := url.Values{}
		q.Set("band", band)
		q.Set("num", strconv.Itoa(num))
		return a.get(a.api("tuner/clearPreset"), q)
	case "auto-preset":
		if len(args) < 2 {
			return fmt.Errorf("tuner auto-preset: missing action")
		}
		switch strings.ToLower(args[1]) {
		case "start":
			q := url.Values{}
			q.Set("band", "fm")
			return a.get(a.api("tuner/startAutoPreset"), q)
		case "cancel":
			return a.get(a.api("tuner/cancelAutoPreset"), nil)
		default:
			return fmt.Errorf("tuner auto-preset: invalid action %s", args[1])
		}
	case "dab-scan":
		if len(args) < 2 {
			return fmt.Errorf("tuner dab-scan: missing action")
		}
		switch strings.ToLower(args[1]) {
		case "start":
			return a.get(a.api("tuner/startDabInitialScan"), nil)
		case "cancel":
			return a.get(a.api("tuner/cancelDabInitialScan"), nil)
		default:
			return fmt.Errorf("tuner dab-scan: invalid action %s", args[1])
		}
	case "dab-service":
		dir, err := cmd.Flags().GetString("dir")
		if err != nil {
			return err
		}
		q := url.Values{}
		q.Set("dir", dir)
		return a.get(a.api("tuner/setDabService"), q)
	default:
		return fmt.Errorf("tuner: unknown command %s", args[0])
	}
}

func zoneOrDefault(zone string) string {
	zone = strings.TrimSpace(zone)
	if zone == "" {
		return "main"
	}
	return zone
}

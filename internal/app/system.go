package app

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

func (a *App) System(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("system: missing subcommand")
	}
	switch args[0] {
	case "speaker-a":
		enable, err := cmd.Flags().GetBool("enable")
		if err != nil {
			return err
		}
		q := url.Values{}
		q.Set("enable", strconv.FormatBool(enable))
		return a.get(a.api("system/setSpeakerA"), q)
	case "speaker-b":
		enable, err := cmd.Flags().GetBool("enable")
		if err != nil {
			return err
		}
		q := url.Values{}
		q.Set("enable", strconv.FormatBool(enable))
		return a.get(a.api("system/setSpeakerB"), q)
	case "dimmer":
		value, err := cmd.Flags().GetInt("value")
		if err != nil {
			return err
		}
		q := url.Values{}
		q.Set("value", strconv.Itoa(value))
		return a.get(a.api("system/setDimmer"), q)
	case "zoneb-volume-sync":
		enable, err := cmd.Flags().GetBool("enable")
		if err != nil {
			return err
		}
		q := url.Values{}
		q.Set("enable", strconv.FormatBool(enable))
		return a.get(a.api("system/setZoneBVolumeSync"), q)
	case "hdmi-out":
		if len(args) < 2 {
			return fmt.Errorf("system hdmi-out: missing output number")
		}
		enable, err := cmd.Flags().GetBool("enable")
		if err != nil {
			return err
		}
		out := strings.TrimSpace(args[1])
		q := url.Values{}
		q.Set("enable", strconv.FormatBool(enable))
		switch out {
		case "1":
			return a.get(a.api("system/setHdmiOut1"), q)
		case "2":
			return a.get(a.api("system/setHdmiOut2"), q)
		default:
			return fmt.Errorf("system hdmi-out: invalid output %s", out)
		}
	case "name":
		if len(args) < 2 {
			return fmt.Errorf("system name: missing action")
		}
		switch args[1] {
		case "get":
			id, err := cmd.Flags().GetString("id")
			if err != nil {
				return err
			}
			q := url.Values{}
			q.Set("id", id)
			return a.get(a.api("system/getNameText"), q)
		case "set":
			id, err := cmd.Flags().GetString("id")
			if err != nil {
				return err
			}
			text, err := cmd.Flags().GetString("text")
			if err != nil {
				return err
			}
			q := url.Values{}
			q.Set("id", id)
			q.Set("text", text)
			return a.get(a.api("system/setNameText"), q)
		default:
			return fmt.Errorf("system name: unknown action %s", args[1])
		}
	case "location":
		return a.get(a.api("system/getLocationInfo"), nil)
	case "ir":
		code, err := cmd.Flags().GetString("code")
		if err != nil {
			return err
		}
		q := url.Values{}
		q.Set("code", code)
		return a.get(a.api("system/sendIrCode"), q)
	case "auto-play":
		enable, err := cmd.Flags().GetBool("enable")
		if err != nil {
			return err
		}
		q := url.Values{}
		q.Set("enable", strconv.FormatBool(enable))
		return a.get(a.api("system/setAutoPlay"), q)
	case "speaker-pattern":
		num, err := cmd.Flags().GetInt("num")
		if err != nil {
			return err
		}
		q := url.Values{}
		q.Set("num", strconv.Itoa(num))
		return a.get(a.api("system/setSpeakerPattern"), q)
	case "party-mode":
		enable, err := cmd.Flags().GetBool("enable")
		if err != nil {
			return err
		}
		q := url.Values{}
		q.Set("enable", strconv.FormatBool(enable))
		return a.get(a.api("system/setPartyMode"), q)
	case "reboot":
		scope, err := cmd.Flags().GetString("scope")
		if err != nil {
			return err
		}
		switch strings.ToLower(strings.TrimSpace(scope)) {
		case "network":
			return a.get(a.api("system/requestNetworkReboot"), nil)
		case "system":
			return a.get(a.api("system/requestSystemReboot"), nil)
		default:
			return fmt.Errorf("system reboot: invalid scope %s", scope)
		}
	default:
		return fmt.Errorf("system: unknown command %s", args[0])
	}
}

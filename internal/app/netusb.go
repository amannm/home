package app

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

func (a *App) Netusb(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("netusb: missing subcommand")
	}
	switch args[0] {
	case "preset-info":
		return a.get(a.api("netusb/getPresetInfo"), nil)
	case "play-info":
		return a.get(a.api("netusb/getPlayInfo"), nil)
	case "playback":
		if len(args) < 2 {
			return fmt.Errorf("netusb playback: missing value")
		}
		q := url.Values{}
		q.Set("playback", args[1])
		return a.get(a.api("netusb/setPlayback"), q)
	case "seek":
		position, err := cmd.Flags().GetInt("position")
		if err != nil {
			return err
		}
		q := url.Values{}
		q.Set("position", strconv.Itoa(position))
		return a.get(a.api("netusb/setPlayPosition"), q)
	case "repeat":
		if len(args) < 2 {
			return fmt.Errorf("netusb repeat: missing value")
		}
		q := url.Values{}
		q.Set("mode", args[1])
		return a.get(a.api("netusb/setRepeat"), q)
	case "shuffle":
		if len(args) < 2 {
			return fmt.Errorf("netusb shuffle: missing value")
		}
		q := url.Values{}
		q.Set("mode", args[1])
		return a.get(a.api("netusb/setShuffle"), q)
	case "repeat-toggle":
		return a.get(a.api("netusb/toggleRepeat"), nil)
	case "shuffle-toggle":
		return a.get(a.api("netusb/toggleShuffle"), nil)
	case "list":
		input, err := cmd.Flags().GetString("input")
		if err != nil {
			return err
		}
		index, err := cmd.Flags().GetInt("index")
		if err != nil {
			return err
		}
		size, err := cmd.Flags().GetInt("size")
		if err != nil {
			return err
		}
		lang, err := cmd.Flags().GetString("lang")
		if err != nil {
			return err
		}
		q := url.Values{}
		q.Set("input", input)
		if cmd.Flags().Changed("index") {
			q.Set("index", strconv.Itoa(index))
		}
		if cmd.Flags().Changed("size") {
			q.Set("size", strconv.Itoa(size))
		}
		if strings.TrimSpace(lang) != "" {
			q.Set("lang", lang)
		}
		return a.get(a.api("netusb/getListInfo"), q)
	case "list-control":
		typ, err := cmd.Flags().GetString("type")
		if err != nil {
			return err
		}
		index, err := cmd.Flags().GetInt("index")
		if err != nil {
			return err
		}
		listID, err := cmd.Flags().GetString("list-id")
		if err != nil {
			return err
		}
		q := url.Values{}
		if strings.TrimSpace(listID) != "" {
			q.Set("list_id", listID)
		}
		q.Set("type", typ)
		if cmd.Flags().Changed("index") {
			q.Set("index", strconv.Itoa(index))
		}
		q.Set("zone", zoneOrDefault(a.Options.Zone))
		return a.get(a.api("netusb/setListControl"), q)
	case "search":
		listID, err := cmd.Flags().GetString("list-id")
		if err != nil {
			return err
		}
		query, err := cmd.Flags().GetString("string")
		if err != nil {
			return err
		}
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
			if strings.TrimSpace(query) == "" {
				return fmt.Errorf("netusb search: --string or --stdin is required")
			}
			payload := map[string]any{"string": query}
			if strings.TrimSpace(listID) != "" {
				payload["list_id"] = listID
			}
			body, err = json.Marshal(payload)
			if err != nil {
				return err
			}
		}
		return a.post(a.api("netusb/setSearchString"), nil, body)
	case "preset":
		if len(args) < 2 {
			return fmt.Errorf("netusb preset: missing action")
		}
		switch args[1] {
		case "recall":
			num, err := cmd.Flags().GetInt("num")
			if err != nil {
				return err
			}
			q := url.Values{}
			q.Set("zone", zoneOrDefault(a.Options.Zone))
			q.Set("num", strconv.Itoa(num))
			return a.get(a.api("netusb/recallPreset"), q)
		case "store":
			num, err := cmd.Flags().GetInt("num")
			if err != nil {
				return err
			}
			q := url.Values{}
			q.Set("num", strconv.Itoa(num))
			return a.get(a.api("netusb/storePreset"), q)
		case "clear":
			num, err := cmd.Flags().GetInt("num")
			if err != nil {
				return err
			}
			q := url.Values{}
			q.Set("num", strconv.Itoa(num))
			return a.get(a.api("netusb/clearPreset"), q)
		case "move":
			from, err := cmd.Flags().GetInt("from")
			if err != nil {
				return err
			}
			to, err := cmd.Flags().GetInt("to")
			if err != nil {
				return err
			}
			q := url.Values{}
			q.Set("from", strconv.Itoa(from))
			q.Set("to", strconv.Itoa(to))
			return a.get(a.api("netusb/movePreset"), q)
		default:
			return fmt.Errorf("netusb preset: unknown action %s", args[1])
		}
	case "recent":
		if len(args) < 2 {
			return fmt.Errorf("netusb recent: missing action")
		}
		switch args[1] {
		case "get":
			return a.get(a.api("netusb/getRecentInfo"), nil)
		case "recall":
			num, err := cmd.Flags().GetInt("num")
			if err != nil {
				return err
			}
			q := url.Values{}
			q.Set("zone", zoneOrDefault(a.Options.Zone))
			q.Set("num", strconv.Itoa(num))
			return a.get(a.api("netusb/recallRecentItem"), q)
		case "clear":
			return a.get(a.api("netusb/clearRecentInfo"), nil)
		default:
			return fmt.Errorf("netusb recent: unknown action %s", args[1])
		}
	case "settings":
		return a.get(a.api("netusb/getSettings"), nil)
	case "quality":
		input, err := cmd.Flags().GetString("input")
		if err != nil {
			return err
		}
		value, err := cmd.Flags().GetString("value")
		if err != nil {
			return err
		}
		q := url.Values{}
		q.Set("input", input)
		q.Set("value", value)
		return a.get(a.api("netusb/setQuality"), q)
	case "account-status":
		return a.get(a.api("netusb/getAccountStatus"), nil)
	case "service-info":
		input, err := cmd.Flags().GetString("input")
		if err != nil {
			return err
		}
		typ, err := cmd.Flags().GetString("type")
		if err != nil {
			return err
		}
		q := url.Values{}
		q.Set("input", input)
		if strings.TrimSpace(typ) != "" {
			q.Set("type", typ)
		}
		return a.get(a.api("netusb/getServiceInfo"), q)
	default:
		return fmt.Errorf("netusb: unknown command %s", args[0])
	}
}

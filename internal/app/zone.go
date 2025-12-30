package app

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

func (a *App) Zone(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("zone: missing subcommand")
	}
	zone := strings.TrimSpace(a.Options.Zone)
	if zone == "" {
		zone = "main"
	}
	zpath := func(p string) string {
		return a.api(zone + "/" + p)
	}
	switch args[0] {
	case "status":
		return a.get(zpath("getStatus"), nil)
	case "sound-programs":
		return a.get(zpath("getSoundProgramList"), nil)
	case "power":
		if len(args) < 2 {
			return fmt.Errorf("zone power: missing value")
		}
		q := url.Values{}
		q.Set("power", args[1])
		return a.get(zpath("setPower"), q)
	case "sleep":
		if len(args) < 2 {
			return fmt.Errorf("zone sleep: missing value")
		}
		q := url.Values{}
		q.Set("sleep", args[1])
		return a.get(zpath("setSleep"), q)
	case "volume":
		if len(args) < 2 {
			return fmt.Errorf("zone volume: missing value")
		}
		step, err := cmd.Flags().GetInt("step")
		if err != nil {
			return err
		}
		q := url.Values{}
		q.Set("volume", args[1])
		if args[1] == "up" || args[1] == "down" {
			q.Set("step", strconv.Itoa(step))
		}
		return a.get(zpath("setVolume"), q)
	case "mute":
		enable, err := cmd.Flags().GetBool("enable")
		if err != nil {
			return err
		}
		q := url.Values{}
		q.Set("enable", strconv.FormatBool(enable))
		return a.get(zpath("setMute"), q)
	case "input":
		if len(args) < 2 {
			return fmt.Errorf("zone input: missing input id")
		}
		mode, err := cmd.Flags().GetString("mode")
		if err != nil {
			return err
		}
		q := url.Values{}
		q.Set("input", args[1])
		if strings.TrimSpace(mode) != "" {
			q.Set("mode", mode)
		}
		return a.get(zpath("setInput"), q)
	case "sound-program":
		if len(args) < 2 {
			return fmt.Errorf("zone sound-program: missing id")
		}
		q := url.Values{}
		q.Set("program", args[1])
		return a.get(zpath("setSoundProgram"), q)
	case "surround-3d":
		enable, err := cmd.Flags().GetBool("enable")
		if err != nil {
			return err
		}
		q := url.Values{}
		q.Set("enable", strconv.FormatBool(enable))
		return a.get(zpath("set3dSurround"), q)
	case "direct":
		enable, err := cmd.Flags().GetBool("enable")
		if err != nil {
			return err
		}
		q := url.Values{}
		q.Set("enable", strconv.FormatBool(enable))
		return a.get(zpath("setDirect"), q)
	case "pure-direct":
		enable, err := cmd.Flags().GetBool("enable")
		if err != nil {
			return err
		}
		q := url.Values{}
		q.Set("enable", strconv.FormatBool(enable))
		return a.get(zpath("setPureDirect"), q)
	case "enhancer":
		enable, err := cmd.Flags().GetBool("enable")
		if err != nil {
			return err
		}
		q := url.Values{}
		q.Set("enable", strconv.FormatBool(enable))
		return a.get(zpath("setEnhancer"), q)
	case "tone":
		mode, err := cmd.Flags().GetString("mode")
		if err != nil {
			return err
		}
		bass, err := cmd.Flags().GetInt("bass")
		if err != nil {
			return err
		}
		treble, err := cmd.Flags().GetInt("treble")
		if err != nil {
			return err
		}
		q := url.Values{}
		if strings.TrimSpace(mode) != "" {
			q.Set("mode", mode)
		}
		if cmd.Flags().Changed("bass") {
			q.Set("bass", strconv.Itoa(bass))
		}
		if cmd.Flags().Changed("treble") {
			q.Set("treble", strconv.Itoa(treble))
		}
		return a.get(zpath("setToneControl"), q)
	case "eq":
		mode, err := cmd.Flags().GetString("mode")
		if err != nil {
			return err
		}
		low, err := cmd.Flags().GetInt("low")
		if err != nil {
			return err
		}
		mid, err := cmd.Flags().GetInt("mid")
		if err != nil {
			return err
		}
		high, err := cmd.Flags().GetInt("high")
		if err != nil {
			return err
		}
		q := url.Values{}
		if strings.TrimSpace(mode) != "" {
			q.Set("mode", mode)
		}
		if cmd.Flags().Changed("low") {
			q.Set("low", strconv.Itoa(low))
		}
		if cmd.Flags().Changed("mid") {
			q.Set("mid", strconv.Itoa(mid))
		}
		if cmd.Flags().Changed("high") {
			q.Set("high", strconv.Itoa(high))
		}
		return a.get(zpath("setEqualizer"), q)
	case "balance":
		value, err := cmd.Flags().GetInt("value")
		if err != nil {
			return err
		}
		q := url.Values{}
		q.Set("value", strconv.Itoa(value))
		return a.get(zpath("setBalance"), q)
	case "dialogue-level":
		value, err := cmd.Flags().GetInt("value")
		if err != nil {
			return err
		}
		q := url.Values{}
		q.Set("value", strconv.Itoa(value))
		return a.get(zpath("setDialogueLevel"), q)
	case "dialogue-lift":
		value, err := cmd.Flags().GetInt("value")
		if err != nil {
			return err
		}
		q := url.Values{}
		q.Set("value", strconv.Itoa(value))
		return a.get(zpath("setDialogueLift"), q)
	case "clear-voice":
		enable, err := cmd.Flags().GetBool("enable")
		if err != nil {
			return err
		}
		q := url.Values{}
		q.Set("enable", strconv.FormatBool(enable))
		return a.get(zpath("setClearVoice"), q)
	case "subwoofer-volume":
		volume, err := cmd.Flags().GetInt("volume")
		if err != nil {
			return err
		}
		q := url.Values{}
		q.Set("volume", strconv.Itoa(volume))
		return a.get(zpath("setSubwooferVolume"), q)
	case "bass-extension":
		enable, err := cmd.Flags().GetBool("enable")
		if err != nil {
			return err
		}
		q := url.Values{}
		q.Set("enable", strconv.FormatBool(enable))
		return a.get(zpath("setBassExtension"), q)
	case "signal":
		return a.get(zpath("getSignalInfo"), nil)
	case "prepare-input":
		input, err := cmd.Flags().GetString("input")
		if err != nil {
			return err
		}
		q := url.Values{}
		q.Set("input", input)
		return a.get(zpath("prepareInputChange"), q)
	case "scene":
		num, err := cmd.Flags().GetInt("num")
		if err != nil {
			return err
		}
		q := url.Values{}
		q.Set("num", strconv.Itoa(num))
		return a.get(zpath("recallScene"), q)
	case "osd":
		enable, err := cmd.Flags().GetBool("enable")
		if err != nil {
			return err
		}
		q := url.Values{}
		q.Set("enable", strconv.FormatBool(enable))
		return a.get(zpath("setContentsDisplay"), q)
	case "cursor":
		if len(args) < 2 {
			return fmt.Errorf("zone cursor: missing value")
		}
		q := url.Values{}
		q.Set("cursor", args[1])
		return a.get(zpath("controlCursor"), q)
	case "menu":
		if len(args) < 2 {
			return fmt.Errorf("zone menu: missing value")
		}
		q := url.Values{}
		q.Set("menu", args[1])
		return a.get(zpath("executeMenu"), q)
	case "actual-volume":
		mode, err := cmd.Flags().GetString("mode")
		if err != nil {
			return err
		}
		value, err := cmd.Flags().GetFloat64("value")
		if err != nil {
			return err
		}
		q := url.Values{}
		q.Set("mode", mode)
		if cmd.Flags().Changed("value") {
			q.Set("value", strconv.FormatFloat(value, 'f', -1, 64))
		}
		return a.get(zpath("setActualVolume"), q)
	case "surround-decoder":
		typ, err := cmd.Flags().GetString("type")
		if err != nil {
			return err
		}
		q := url.Values{}
		q.Set("type", typ)
		return a.get(zpath("setSurroundDecoderType"), q)
	case "link-control":
		control, err := cmd.Flags().GetString("control")
		if err != nil {
			return err
		}
		q := url.Values{}
		q.Set("control", control)
		return a.get(zpath("setLinkControl"), q)
	case "link-delay":
		delay, err := cmd.Flags().GetString("delay")
		if err != nil {
			return err
		}
		q := url.Values{}
		q.Set("delay", delay)
		return a.get(zpath("setLinkAudioDelay"), q)
	case "link-quality":
		quality, err := cmd.Flags().GetString("quality")
		if err != nil {
			return err
		}
		q := url.Values{}
		q.Set("quality", quality)
		return a.get(zpath("setLinkAudioQuality"), q)
	default:
		return fmt.Errorf("zone: unknown command %s", args[0])
	}
}

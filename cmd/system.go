package cmd

import (
	"github.com/amannm/yxc/internal/app"
	"github.com/spf13/cobra"
)

func runSystem(prefix ...string) func(*cobra.Command, []string) error {
	return func(cmd *cobra.Command, args []string) error {
		return app.New(opts).System(cmd, append(prefix, args...))
	}
}

func newSystemCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "system",
		Short: "System-level controls",
	}

	cmd.AddCommand(
		newSystemSpeakerACmd(),
		newSystemSpeakerBCmd(),
		newSystemDimmerCmd(),
		newSystemZoneBVolumeSyncCmd(),
		newSystemHdmiOutCmd(),
		newSystemNameCmd(),
		newSystemLocationCmd(),
		newSystemIrCmd(),
		newSystemAutoPlayCmd(),
		newSystemSpeakerPatternCmd(),
		newSystemPartyModeCmd(),
		newSystemRebootCmd(),
	)

	return cmd
}

func newSystemSpeakerACmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "speaker-a",
		Short: "Enable/disable Speaker A output",
		RunE:  runSystem("speaker-a"),
	}

	cmd.Flags().Bool("enable", false, "Enable/disable Speaker A output")
	_ = cmd.MarkFlagRequired("enable")

	return cmd
}

func newSystemSpeakerBCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "speaker-b",
		Short: "Enable/disable Speaker B output",
		RunE:  runSystem("speaker-b"),
	}

	cmd.Flags().Bool("enable", false, "Enable/disable Speaker B output")
	_ = cmd.MarkFlagRequired("enable")

	return cmd
}

func newSystemDimmerCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "dimmer",
		Short: "Set display dimmer",
		RunE:  runSystem("dimmer"),
	}

	cmd.Flags().Int("value", 0, "Dimmer value (-1 for auto, 0+ for manual)")
	_ = cmd.MarkFlagRequired("value")

	return cmd
}

func newSystemZoneBVolumeSyncCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "zoneb-volume-sync",
		Short: "Enable/disable Zone B volume synchronization",
		RunE:  runSystem("zoneb-volume-sync"),
	}

	cmd.Flags().Bool("enable", false, "Enable/disable Zone B volume synchronization")
	_ = cmd.MarkFlagRequired("enable")

	return cmd
}

func newSystemHdmiOutCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "hdmi-out <1|2>",
		Short: "Enable/disable HDMI outputs",
		Args:  cobra.ExactArgs(1),
		RunE:  runSystem("hdmi-out"),
	}

	cmd.Flags().Bool("enable", false, "Enable/disable HDMI output")
	_ = cmd.MarkFlagRequired("enable")

	return cmd
}

func newSystemNameCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "name",
		Short: "Get/set custom name text for inputs/sound programs",
	}

	cmd.AddCommand(
		newSystemNameGetCmd(),
		newSystemNameSetCmd(),
	)

	return cmd
}

func newSystemNameGetCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get",
		Short: "Get custom name text",
		RunE:  runSystem("name", "get"),
	}

	cmd.Flags().String("id", "", "Input or sound program ID")
	_ = cmd.MarkFlagRequired("id")

	return cmd
}

func newSystemNameSetCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set",
		Short: "Set custom name text",
		RunE:  runSystem("name", "set"),
	}

	cmd.Flags().String("id", "", "Input or sound program ID")
	cmd.Flags().String("text", "", "New text (max 64 chars)")
	_ = cmd.MarkFlagRequired("id")
	_ = cmd.MarkFlagRequired("text")

	return cmd
}

func newSystemLocationCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "location",
		Short: "Get MusicCast network location info",
		RunE:  runSystem("location"),
	}

	return cmd
}

func newSystemIrCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ir",
		Short: "Send IR remote code",
		RunE:  runSystem("ir"),
	}

	cmd.Flags().String("code", "", "IR code (hex string)")
	_ = cmd.MarkFlagRequired("code")

	return cmd
}

func newSystemAutoPlayCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "auto-play",
		Short: "Configure auto play behavior",
		RunE:  runSystem("auto-play"),
	}

	cmd.Flags().Bool("enable", false, "Enable/disable auto play")
	_ = cmd.MarkFlagRequired("enable")

	return cmd
}

func newSystemSpeakerPatternCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "speaker-pattern",
		Short: "Set speaker configuration pattern",
		RunE:  runSystem("speaker-pattern"),
	}

	cmd.Flags().Int("num", 0, "Speaker configuration pattern number")
	_ = cmd.MarkFlagRequired("num")

	return cmd
}

func newSystemPartyModeCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "party-mode",
		Short: "Enable/disable party mode",
		RunE:  runSystem("party-mode"),
	}

	cmd.Flags().Bool("enable", false, "Enable/disable party mode")
	_ = cmd.MarkFlagRequired("enable")

	return cmd
}

func newSystemRebootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "reboot",
		Short: "Reboot network module or entire system",
		RunE:  runSystem("reboot"),
	}

	cmd.Flags().String("scope", "", "Reboot scope: network|system")
	_ = cmd.MarkFlagRequired("scope")

	return cmd
}

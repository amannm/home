package cmd

import (
	"github.com/amannm/yxc/internal/app"
	"github.com/spf13/cobra"
)

func runZone(prefix ...string) func(*cobra.Command, []string) error {
	return func(cmd *cobra.Command, args []string) error {
		return app.New(opts).Zone(cmd, append(prefix, args...))
	}
}

func newZoneCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "zone",
		Short: "Zone controls",
	}

	cmd.AddCommand(
		newZoneStatusCmd(),
		newZoneSoundProgramsCmd(),
		newZonePowerCmd(),
		newZoneSleepCmd(),
		newZoneVolumeCmd(),
		newZoneMuteCmd(),
		newZoneInputCmd(),
		newZoneSoundProgramCmd(),
		newZoneSurround3DCmd(),
		newZoneDirectCmd(),
		newZonePureDirectCmd(),
		newZoneEnhancerCmd(),
		newZoneToneCmd(),
		newZoneEqCmd(),
		newZoneBalanceCmd(),
		newZoneDialogueLevelCmd(),
		newZoneDialogueLiftCmd(),
		newZoneClearVoiceCmd(),
		newZoneSubwooferVolumeCmd(),
		newZoneBassExtensionCmd(),
		newZoneSignalCmd(),
		newZonePrepareInputCmd(),
		newZoneSceneCmd(),
		newZoneOSDCmd(),
		newZoneCursorCmd(),
		newZoneMenuCmd(),
		newZoneActualVolumeCmd(),
		newZoneSurroundDecoderCmd(),
		newZoneLinkControlCmd(),
		newZoneLinkDelayCmd(),
		newZoneLinkQualityCmd(),
	)

	return cmd
}

func newZoneStatusCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "status",
		Short: "Get zone status",
		RunE:  runZone("status"),
	}

	return cmd
}

func newZoneSoundProgramsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sound-programs",
		Short: "List available sound programs",
		RunE:  runZone("sound-programs"),
	}

	return cmd
}

func newZonePowerCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "power <on|standby|toggle>",
		Short: "Set zone power",
		Args:  cobra.ExactArgs(1),
		RunE:  runZone("power"),
	}

	return cmd
}

func newZoneSleepCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sleep <0|30|60|90|120>",
		Short: "Set sleep timer",
		Args:  cobra.ExactArgs(1),
		RunE:  runZone("sleep"),
	}

	return cmd
}

func newZoneVolumeCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "volume <int|up|down>",
		Short: "Set volume or issue up/down",
		Args:  cobra.ExactArgs(1),
		RunE:  runZone("volume"),
	}

	cmd.Flags().Int("step", 1, "Step count for up/down")

	return cmd
}

func newZoneMuteCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "mute",
		Short: "Enable/disable mute",
		RunE:  runZone("mute"),
	}

	cmd.Flags().Bool("enable", false, "Enable/disable mute")
	_ = cmd.MarkFlagRequired("enable")

	return cmd
}

func newZoneInputCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "input <input-id>",
		Short: "Set input",
		Args:  cobra.ExactArgs(1),
		RunE:  runZone("input"),
	}

	cmd.Flags().String("mode", "", "Input mode (e.g. autoplay_disabled)")

	return cmd
}

func newZoneSoundProgramCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sound-program <program-id>",
		Short: "Set sound program",
		Args:  cobra.ExactArgs(1),
		RunE:  runZone("sound-program"),
	}

	return cmd
}

func newZoneSurround3DCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "surround-3d",
		Short: "Enable/disable 3D surround",
		RunE:  runZone("surround-3d"),
	}

	cmd.Flags().Bool("enable", false, "Enable/disable 3D surround")
	_ = cmd.MarkFlagRequired("enable")

	return cmd
}

func newZoneDirectCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "direct",
		Short: "Enable/disable direct",
		RunE:  runZone("direct"),
	}

	cmd.Flags().Bool("enable", false, "Enable/disable direct")
	_ = cmd.MarkFlagRequired("enable")

	return cmd
}

func newZonePureDirectCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "pure-direct",
		Short: "Enable/disable pure direct",
		RunE:  runZone("pure-direct"),
	}

	cmd.Flags().Bool("enable", false, "Enable/disable pure direct")
	_ = cmd.MarkFlagRequired("enable")

	return cmd
}

func newZoneEnhancerCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "enhancer",
		Short: "Enable/disable compressed music enhancer",
		RunE:  runZone("enhancer"),
	}

	cmd.Flags().Bool("enable", false, "Enable/disable enhancer")
	_ = cmd.MarkFlagRequired("enable")

	return cmd
}

func newZoneToneCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "tone",
		Short: "Set tone control",
		RunE:  runZone("tone"),
	}

	cmd.Flags().String("mode", "", "Tone mode: manual|auto|bypass")
	cmd.Flags().Int("bass", 0, "Bass level")
	cmd.Flags().Int("treble", 0, "Treble level")

	return cmd
}

func newZoneEqCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "eq",
		Short: "Set equalizer",
		RunE:  runZone("eq"),
	}

	cmd.Flags().String("mode", "", "EQ mode: manual|auto|bypass")
	cmd.Flags().Int("low", 0, "Low band")
	cmd.Flags().Int("mid", 0, "Mid band")
	cmd.Flags().Int("high", 0, "High band")

	return cmd
}

func newZoneBalanceCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "balance",
		Short: "Set L/R balance",
		RunE:  runZone("balance"),
	}

	cmd.Flags().Int("value", 0, "Balance value (negative=left, positive=right)")
	_ = cmd.MarkFlagRequired("value")

	return cmd
}

func newZoneDialogueLevelCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "dialogue-level",
		Short: "Set dialogue/vocal enhancement level",
		RunE:  runZone("dialogue-level"),
	}

	cmd.Flags().Int("value", 0, "Dialogue level")
	_ = cmd.MarkFlagRequired("value")

	return cmd
}

func newZoneDialogueLiftCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "dialogue-lift",
		Short: "Set dialogue lift height",
		RunE:  runZone("dialogue-lift"),
	}

	cmd.Flags().Int("value", 0, "Dialogue lift height")
	_ = cmd.MarkFlagRequired("value")

	return cmd
}

func newZoneClearVoiceCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "clear-voice",
		Short: "Enable/disable clear voice",
		RunE:  runZone("clear-voice"),
	}

	cmd.Flags().Bool("enable", false, "Enable/disable clear voice")
	_ = cmd.MarkFlagRequired("enable")

	return cmd
}

func newZoneSubwooferVolumeCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "subwoofer-volume",
		Short: "Set subwoofer volume",
		RunE:  runZone("subwoofer-volume"),
	}

	cmd.Flags().Int("volume", 0, "Subwoofer volume")
	_ = cmd.MarkFlagRequired("volume")

	return cmd
}

func newZoneBassExtensionCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "bass-extension",
		Short: "Enable/disable bass extension",
		RunE:  runZone("bass-extension"),
	}

	cmd.Flags().Bool("enable", false, "Enable/disable bass extension")
	_ = cmd.MarkFlagRequired("enable")

	return cmd
}

func newZoneSignalCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "signal",
		Short: "Get audio signal info",
		RunE:  runZone("signal"),
	}

	return cmd
}

func newZonePrepareInputCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "prepare-input",
		Short: "Prepare for input change",
		RunE:  runZone("prepare-input"),
	}

	cmd.Flags().String("input", "", "Input ID")
	_ = cmd.MarkFlagRequired("input")

	return cmd
}

func newZoneSceneCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "scene",
		Short: "Recall a scene",
		RunE:  runZone("scene"),
	}

	cmd.Flags().Int("num", 0, "Scene number")
	_ = cmd.MarkFlagRequired("num")

	return cmd
}

func newZoneOSDCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "osd",
		Short: "Enable/disable on-screen display",
		RunE:  runZone("osd"),
	}

	cmd.Flags().Bool("enable", false, "Enable/disable OSD")
	_ = cmd.MarkFlagRequired("enable")

	return cmd
}

func newZoneCursorCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cursor <up|down|left|right|select|return>",
		Short: "Send cursor command",
		Args:  cobra.ExactArgs(1),
		RunE:  runZone("cursor"),
	}

	return cmd
}

func newZoneMenuCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "menu <command>",
		Short: "Execute menu command",
		Args:  cobra.ExactArgs(1),
		RunE:  runZone("menu"),
	}

	return cmd
}

func newZoneActualVolumeCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "actual-volume",
		Short: "Set volume display mode and optional value",
		RunE:  runZone("actual-volume"),
	}

	cmd.Flags().String("mode", "", "Mode: db|numeric")
	cmd.Flags().Float64("value", 0, "Numeric value")
	_ = cmd.MarkFlagRequired("mode")

	return cmd
}

func newZoneSurroundDecoderCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "surround-decoder",
		Short: "Set surround decoder type",
		RunE:  runZone("surround-decoder"),
	}

	cmd.Flags().String("type", "", "Decoder type")
	_ = cmd.MarkFlagRequired("type")

	return cmd
}

func newZoneLinkControlCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "link-control",
		Short: "Set MusicCast Link control mode",
		RunE:  runZone("link-control"),
	}

	cmd.Flags().String("control", "", "Control mode: standard|stability|speed")
	_ = cmd.MarkFlagRequired("control")

	return cmd
}

func newZoneLinkDelayCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "link-delay",
		Short: "Set MusicCast Link audio delay mode",
		RunE:  runZone("link-delay"),
	}

	cmd.Flags().String("delay", "", "Delay mode")
	_ = cmd.MarkFlagRequired("delay")

	return cmd
}

func newZoneLinkQualityCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "link-quality",
		Short: "Set MusicCast Link audio quality",
		RunE:  runZone("link-quality"),
	}

	cmd.Flags().String("quality", "", "Quality: compressed|uncompressed")
	_ = cmd.MarkFlagRequired("quality")

	return cmd
}

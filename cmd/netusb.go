package cmd

import (
	"github.com/amannm/yxc/internal/app"
	"github.com/spf13/cobra"
)

func runNetusb(prefix ...string) func(*cobra.Command, []string) error {
	return func(cmd *cobra.Command, args []string) error {
		return app.New(opts).Netusb(cmd, append(prefix, args...))
	}
}

func newNetusbCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "netusb",
		Short: "Network/USB playback and browsing",
	}

	cmd.AddCommand(
		newNetusbPresetInfoCmd(),
		newNetusbPlayInfoCmd(),
		newNetusbPlaybackCmd(),
		newNetusbSeekCmd(),
		newNetusbRepeatCmd(),
		newNetusbShuffleCmd(),
		newNetusbRepeatToggleCmd(),
		newNetusbShuffleToggleCmd(),
		newNetusbListCmd(),
		newNetusbListControlCmd(),
		newNetusbSearchCmd(),
		newNetusbPresetCmd(),
		newNetusbRecentCmd(),
		newNetusbSettingsCmd(),
		newNetusbQualityCmd(),
		newNetusbAccountStatusCmd(),
		newNetusbServiceInfoCmd(),
	)

	return cmd
}

func newNetusbPresetInfoCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "preset-info",
		Short: "Get Net/USB preset info",
		RunE:  runNetusb("preset-info"),
	}

	return cmd
}

func newNetusbPlayInfoCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "play-info",
		Short: "Get current Net/USB play info",
		RunE:  runNetusb("play-info"),
	}

	return cmd
}

func newNetusbPlaybackCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "playback <play|stop|pause|play_pause|previous|next|fast_reverse_start|fast_reverse_end|fast_forward_start|fast_forward_end>",
		Short: "Control playback state",
		Args:  cobra.ExactArgs(1),
		RunE:  runNetusb("playback"),
	}

	return cmd
}

func newNetusbSeekCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "seek",
		Short: "Set play position (seconds)",
		RunE:  runNetusb("seek"),
	}

	cmd.Flags().Int("position", 0, "Position in seconds")
	_ = cmd.MarkFlagRequired("position")

	return cmd
}

func newNetusbRepeatCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "repeat <off|one|all>",
		Short: "Set repeat mode",
		Args:  cobra.ExactArgs(1),
		RunE:  runNetusb("repeat"),
	}

	return cmd
}

func newNetusbShuffleCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "shuffle <off|on|songs|albums>",
		Short: "Set shuffle mode",
		Args:  cobra.ExactArgs(1),
		RunE:  runNetusb("shuffle"),
	}

	return cmd
}

func newNetusbRepeatToggleCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "repeat-toggle",
		Short: "Toggle repeat mode",
		RunE:  runNetusb("repeat-toggle"),
	}

	return cmd
}

func newNetusbShuffleToggleCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "shuffle-toggle",
		Short: "Toggle shuffle mode",
		RunE:  runNetusb("shuffle-toggle"),
	}

	return cmd
}

func newNetusbListCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "Get browsable list info",
		RunE:  runNetusb("list"),
	}

	cmd.Flags().String("input", "", "Input ID")
	cmd.Flags().Int("index", 0, "List index")
	cmd.Flags().Int("size", 8, "List size (1-8)")
	cmd.Flags().String("lang", "", "Language code")
	_ = cmd.MarkFlagRequired("input")

	return cmd
}

func newNetusbListControlCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-control",
		Short: "Navigate/select/play in list",
		RunE:  runNetusb("list-control"),
	}

	cmd.Flags().String("type", "", "Control type: select|play|return")
	cmd.Flags().Int("index", 0, "List index for select/play")
	cmd.Flags().String("list-id", "main", "List ID")
	_ = cmd.MarkFlagRequired("type")

	return cmd
}

func newNetusbSearchCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "search",
		Short: "Search Net/USB",
		RunE:  runNetusb("search"),
	}

	cmd.Flags().String("list-id", "main", "List ID")
	cmd.Flags().String("string", "", "Search query string")
	cmd.Flags().Bool("stdin", false, "Read JSON body from stdin")

	return cmd
}

func newNetusbPresetCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "preset",
		Short: "Recall/store/clear/move presets",
	}

	cmd.AddCommand(
		newNetusbPresetRecallCmd(),
		newNetusbPresetStoreCmd(),
		newNetusbPresetClearCmd(),
		newNetusbPresetMoveCmd(),
	)

	return cmd
}

func newNetusbPresetRecallCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "recall",
		Short: "Recall preset",
		RunE:  runNetusb("preset", "recall"),
	}

	cmd.Flags().Int("num", 0, "Preset number")
	_ = cmd.MarkFlagRequired("num")

	return cmd
}

func newNetusbPresetStoreCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "store",
		Short: "Store preset",
		RunE:  runNetusb("preset", "store"),
	}

	cmd.Flags().Int("num", 0, "Preset number")
	_ = cmd.MarkFlagRequired("num")

	return cmd
}

func newNetusbPresetClearCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "clear",
		Short: "Clear preset",
		RunE:  runNetusb("preset", "clear"),
	}

	cmd.Flags().Int("num", 0, "Preset number")
	_ = cmd.MarkFlagRequired("num")

	return cmd
}

func newNetusbPresetMoveCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "move",
		Short: "Move preset",
		RunE:  runNetusb("preset", "move"),
	}

	cmd.Flags().Int("from", 0, "Source preset number")
	cmd.Flags().Int("to", 0, "Destination preset number")
	_ = cmd.MarkFlagRequired("from")
	_ = cmd.MarkFlagRequired("to")

	return cmd
}

func newNetusbRecentCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "recent",
		Short: "Get/recall/clear recently played",
	}

	cmd.AddCommand(
		newNetusbRecentGetCmd(),
		newNetusbRecentRecallCmd(),
		newNetusbRecentClearCmd(),
	)

	return cmd
}

func newNetusbRecentGetCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get",
		Short: "Get recent items",
		RunE:  runNetusb("recent", "get"),
	}

	return cmd
}

func newNetusbRecentRecallCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "recall",
		Short: "Recall recent item",
		RunE:  runNetusb("recent", "recall"),
	}

	cmd.Flags().Int("num", 0, "Item number")
	_ = cmd.MarkFlagRequired("num")

	return cmd
}

func newNetusbRecentClearCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "clear",
		Short: "Clear recent items",
		RunE:  runNetusb("recent", "clear"),
	}

	return cmd
}

func newNetusbSettingsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "settings",
		Short: "Get streaming service settings",
		RunE:  runNetusb("settings"),
	}

	return cmd
}

func newNetusbQualityCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "quality",
		Short: "Set streaming quality",
		RunE:  runNetusb("quality"),
	}

	cmd.Flags().String("input", "", "Input ID")
	cmd.Flags().String("value", "", "Quality value")
	_ = cmd.MarkFlagRequired("input")
	_ = cmd.MarkFlagRequired("value")

	return cmd
}

func newNetusbAccountStatusCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "account-status",
		Short: "Get account status",
		RunE:  runNetusb("account-status"),
	}

	return cmd
}

func newNetusbServiceInfoCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "service-info",
		Short: "Get service-specific info",
		RunE:  runNetusb("service-info"),
	}

	cmd.Flags().String("input", "", "Input ID")
	cmd.Flags().String("type", "", "Service info type")
	_ = cmd.MarkFlagRequired("input")

	return cmd
}

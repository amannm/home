package cmd

import (
	"github.com/amannm/home/internal/app"
	"github.com/spf13/cobra"
)

func runCD(prefix ...string) func(*cobra.Command, []string) error {
	return func(cmd *cobra.Command, args []string) error {
		return app.New(opts).CD(cmd, append(prefix, args...))
	}
}

func newCdCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cd",
		Short: "CD playback controls",
	}

	cmd.AddCommand(
		newCdPlayInfoCmd(),
		newCdPlaybackCmd(),
		newCdTrayCmd(),
		newCdRepeatCmd(),
		newCdShuffleCmd(),
		newCdRepeatToggleCmd(),
		newCdShuffleToggleCmd(),
		newCdDirectCmd(),
	)

	return cmd
}

func newCdPlayInfoCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "play-info",
		Short: "Get CD play info",
		RunE:  runCD("play-info"),
	}

	return cmd
}

func newCdPlaybackCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "playback <play|stop|pause|previous|next|fast_reverse_start|fast_reverse_end|fast_forward_start|fast_forward_end|track_select>",
		Short: "Control CD playback",
		Args:  cobra.ExactArgs(1),
		RunE:  runCD("playback"),
	}

	cmd.Flags().Int("num", 0, "Track number for track_select")

	return cmd
}

func newCdTrayCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "tray <toggle>",
		Short: "Toggle CD tray",
		Args:  cobra.ExactArgs(1),
		RunE:  runCD("tray"),
	}

	return cmd
}

func newCdRepeatCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "repeat <off|one|all|folder|a-b>",
		Short: "Set CD repeat",
		Args:  cobra.ExactArgs(1),
		RunE:  runCD("repeat"),
	}

	return cmd
}

func newCdShuffleCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "shuffle <off|on|folder|program>",
		Short: "Set CD shuffle",
		Args:  cobra.ExactArgs(1),
		RunE:  runCD("shuffle"),
	}

	return cmd
}

func newCdRepeatToggleCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "repeat-toggle",
		Short: "Toggle CD repeat",
		RunE:  runCD("repeat-toggle"),
	}

	return cmd
}

func newCdShuffleToggleCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "shuffle-toggle",
		Short: "Toggle CD shuffle",
		RunE:  runCD("shuffle-toggle"),
	}

	return cmd
}

func newCdDirectCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "direct",
		Short: "Enable/disable direct CD playback",
		RunE:  runCD("direct"),
	}

	cmd.Flags().Bool("enable", false, "Enable/disable direct CD playback")
	_ = cmd.MarkFlagRequired("enable")

	return cmd
}

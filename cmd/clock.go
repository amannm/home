package cmd

import (
	"github.com/amannm/yxc/internal/app"
	"github.com/spf13/cobra"
)

func runClock(prefix ...string) func(*cobra.Command, []string) error {
	return func(cmd *cobra.Command, args []string) error {
		return app.New(opts).Clock(cmd, append(prefix, args...))
	}
}

func newClockCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "clock",
		Short: "Clock and alarm settings",
	}

	cmd.AddCommand(
		newClockSettingsCmd(),
		newClockAutoSyncCmd(),
		newClockDatetimeCmd(),
		newClockFormatCmd(),
		newClockAlarmCmd(),
	)

	return cmd
}

func newClockSettingsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "settings",
		Short: "Get clock & alarm settings",
		RunE:  runClock("settings"),
	}

	return cmd
}

func newClockAutoSyncCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "auto-sync",
		Short: "Enable/disable auto sync",
		RunE:  runClock("auto-sync"),
	}

	cmd.Flags().Bool("enable", false, "Enable/disable auto sync")
	_ = cmd.MarkFlagRequired("enable")

	return cmd
}

func newClockDatetimeCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "datetime",
		Short: "Set date/time (YYMMDDhhmmss)",
		RunE:  runClock("datetime"),
	}

	cmd.Flags().String("date-time", "", "Date/time in YYMMDDhhmmss format")
	_ = cmd.MarkFlagRequired("date-time")

	return cmd
}

func newClockFormatCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "format <12h|24h>",
		Short: "Set clock format",
		Args:  cobra.ExactArgs(1),
		RunE:  runClock("format"),
	}

	return cmd
}

func newClockAlarmCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "alarm",
		Short: "Configure alarm settings (POST JSON)",
		RunE:  runClock("alarm"),
	}

	cmd.Flags().String("file", "", "Alarm settings JSON file")
	cmd.Flags().Bool("stdin", false, "Read JSON body from stdin")

	return cmd
}

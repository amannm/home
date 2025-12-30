package cmd

import (
	"github.com/amannm/home/internal/app"
	"github.com/spf13/cobra"
)

func runTuner(prefix ...string) func(*cobra.Command, []string) error {
	return func(cmd *cobra.Command, args []string) error {
		return app.New(opts).Tuner(cmd, append(prefix, args...))
	}
}

func newTunerCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "tuner",
		Short: "Tuner controls",
	}

	cmd.AddCommand(
		newTunerPresetInfoCmd(),
		newTunerPlayInfoCmd(),
		newTunerBandCmd(),
		newTunerFreqCmd(),
		newTunerRecallCmd(),
		newTunerSwitchCmd(),
		newTunerStoreCmd(),
		newTunerClearCmd(),
		newTunerAutoPresetCmd(),
		newTunerDabScanCmd(),
		newTunerDabServiceCmd(),
	)

	return cmd
}

func newTunerPresetInfoCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "preset-info",
		Short: "Get preset info",
		RunE:  runTuner("preset-info"),
	}

	cmd.Flags().String("band", "", "Band: common|am|fm|dab")
	_ = cmd.MarkFlagRequired("band")

	return cmd
}

func newTunerPlayInfoCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "play-info",
		Short: "Get current tuner play info",
		RunE:  runTuner("play-info"),
	}

	return cmd
}

func newTunerBandCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "band <am|fm|dab>",
		Short: "Set band",
		Args:  cobra.ExactArgs(1),
		RunE:  runTuner("band"),
	}

	return cmd
}

func newTunerFreqCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "freq",
		Short: "Set frequency / scan",
		RunE:  runTuner("freq"),
	}

	cmd.Flags().String("band", "", "Band: am|fm")
	cmd.Flags().String("tuning", "", "Tuning: tp_up|tp_down|direct|auto_up|auto_down|cancel")
	cmd.Flags().Int("num", 0, "Frequency value for direct tuning")
	_ = cmd.MarkFlagRequired("band")
	_ = cmd.MarkFlagRequired("tuning")

	return cmd
}

func newTunerRecallCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "recall",
		Short: "Recall preset",
		RunE:  runTuner("recall"),
	}

	cmd.Flags().String("band", "", "Band: common|am|fm|dab")
	cmd.Flags().Int("num", 0, "Preset number")
	_ = cmd.MarkFlagRequired("band")
	_ = cmd.MarkFlagRequired("num")

	return cmd
}

func newTunerSwitchCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "switch",
		Short: "Next/previous preset",
		RunE:  runTuner("switch"),
	}

	cmd.Flags().String("dir", "", "Direction: next|previous")
	_ = cmd.MarkFlagRequired("dir")

	return cmd
}

func newTunerStoreCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "store",
		Short: "Store current station to preset",
		RunE:  runTuner("store"),
	}

	cmd.Flags().Int("num", 0, "Preset number")
	_ = cmd.MarkFlagRequired("num")

	return cmd
}

func newTunerClearCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "clear",
		Short: "Clear preset",
		RunE:  runTuner("clear"),
	}

	cmd.Flags().String("band", "", "Band: common|am|fm|dab")
	cmd.Flags().Int("num", 0, "Preset number")
	_ = cmd.MarkFlagRequired("band")
	_ = cmd.MarkFlagRequired("num")

	return cmd
}

func newTunerAutoPresetCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "auto-preset <start|cancel>",
		Short: "Start/cancel automatic preset scan (FM)",
		Args:  cobra.ExactArgs(1),
		RunE:  runTuner("auto-preset"),
	}

	return cmd
}

func newTunerDabScanCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "dab-scan <start|cancel>",
		Short: "Start/cancel DAB initial scan",
		Args:  cobra.ExactArgs(1),
		RunE:  runTuner("dab-scan"),
	}

	return cmd
}

func newTunerDabServiceCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "dab-service",
		Short: "Next/previous DAB service",
		RunE:  runTuner("dab-service"),
	}

	cmd.Flags().String("dir", "", "Direction: next|previous")
	_ = cmd.MarkFlagRequired("dir")

	return cmd
}

package cmd

import (
	"github.com/amannm/yxc/internal/app"
	"github.com/spf13/cobra"
)

func runDist(prefix ...string) func(*cobra.Command, []string) error {
	return func(cmd *cobra.Command, args []string) error {
		return app.New(opts).Dist(cmd, append(prefix, args...))
	}
}

func newDistCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "dist",
		Short: "MusicCast distribution controls",
	}

	cmd.AddCommand(
		newDistInfoCmd(),
		newDistServerCmd(),
		newDistClientCmd(),
		newDistStartCmd(),
		newDistStopCmd(),
		newDistGroupNameCmd(),
	)

	return cmd
}

func newDistInfoCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "info",
		Short: "Get distribution info",
		RunE:  runDist("info"),
	}

	return cmd
}

func newDistServerCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "server",
		Short: "Configure device as Link server (POST JSON)",
		RunE:  runDist("server"),
	}

	cmd.Flags().String("file", "", "Server configuration JSON file")
	cmd.Flags().Bool("stdin", false, "Read JSON body from stdin")

	return cmd
}

func newDistClientCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "client",
		Short: "Configure device as Link client (POST JSON)",
		RunE:  runDist("client"),
	}

	cmd.Flags().String("file", "", "Client configuration JSON file")
	cmd.Flags().Bool("stdin", false, "Read JSON body from stdin")

	return cmd
}

func newDistStartCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "start",
		Short: "Start distribution",
		RunE:  runDist("start"),
	}

	cmd.Flags().Int("num", 0, "Distribution number")
	_ = cmd.MarkFlagRequired("num")

	return cmd
}

func newDistStopCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "stop",
		Short: "Stop distribution",
		RunE:  runDist("stop"),
	}

	return cmd
}

func newDistGroupNameCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "group-name",
		Short: "Set group name (POST JSON)",
		RunE:  runDist("group-name"),
	}

	cmd.Flags().String("name", "", "Group name")
	cmd.Flags().Bool("stdin", false, "Read JSON body from stdin")

	return cmd
}

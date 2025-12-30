package cmd

import (
	"github.com/amannm/home/internal/app"
	"github.com/spf13/cobra"
)

func newRawCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "raw <method> <path>",
		Short: "Call an arbitrary endpoint by path",
		Args:  cobra.MinimumNArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			return app.New(opts).Raw(cmd, args)
		},
	}

	cmd.Flags().StringArray("query", nil, "Query string pairs (k=v), repeatable")
	cmd.Flags().String("data", "", "JSON request body")
	cmd.Flags().Bool("stdin", false, "Read JSON body from stdin")

	return cmd
}

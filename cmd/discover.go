package cmd

import (
	"github.com/amannm/yxc/internal/app"
	"github.com/spf13/cobra"
)

func newDiscoverCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "discover",
		Short: "Discover Yamaha devices on the local network",
		RunE: func(cmd *cobra.Command, args []string) error {
			return app.New(opts).Discover()
		},
	}

	return cmd
}

package cmd

import (
	"fmt"

	"github.com/amannm/yxc/internal/app"
	"github.com/spf13/cobra"
)

func newVersionCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Print version",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(app.Version)
		},
	}

	return cmd
}

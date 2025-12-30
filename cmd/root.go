package cmd

import (
	"os"
	"time"

	"github.com/amannm/yxc/internal/app"
	"github.com/spf13/cobra"
)

var opts app.Options

var rootCmd = &cobra.Command{
	Use:          "yxc",
	Short:        "CLI for Yamaha Extended Control API",
	Long:         "Interact with Yamaha receivers that expose the Extended Control API over HTTP.",
	SilenceUsage: true,
	Version:      app.Version,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&opts.Host, "host", "H", "", "Receiver host/IP (e.g. 192.168.1.50)")
	rootCmd.PersistentFlags().StringVar(&opts.BaseURL, "base-url", "", "Full base URL (default: http://<host>/YamahaExtendedControl)")
	rootCmd.PersistentFlags().StringVar(&opts.APIPrefix, "api-prefix", "/v1", "API prefix")
	rootCmd.PersistentFlags().StringVar(&opts.Zone, "zone", "main", "Default zone for zone commands: main|zone2|zone3|zone4")
	rootCmd.PersistentFlags().DurationVar(&opts.Timeout, "timeout", 5*time.Second, "Request timeout (e.g. 2s, 500ms)")
	rootCmd.PersistentFlags().IntVar(&opts.Retries, "retries", 0, "Retry count on transient network errors")
	rootCmd.PersistentFlags().StringVar(&opts.Auth, "auth", "", "Basic auth (user:pass)")
	rootCmd.PersistentFlags().StringArrayVar(&opts.Headers, "header", nil, "Add an HTTP header (repeatable)")
	rootCmd.PersistentFlags().BoolVar(&opts.DryRun, "dry-run", false, "Print the HTTP request that would be sent, do not send it")
	rootCmd.PersistentFlags().BoolVar(&opts.Curl, "curl", false, "Output an equivalent curl command")
	rootCmd.PersistentFlags().StringVar(&opts.Format, "format", "pretty", "Output: json|pretty|yaml|table")
	rootCmd.PersistentFlags().StringVar(&opts.JQ, "jq", "", "Apply jq filter to JSON output (requires jq installed)")
	rootCmd.PersistentFlags().CountVarP(&opts.Verbose, "verbose", "v", "Verbose logging (repeatable: -vv for more)")
	rootCmd.PersistentFlags().BoolVarP(&opts.Quiet, "quiet", "q", false, "Only print command output (no status lines)")
	rootCmd.PersistentFlags().BoolVar(&opts.NoColor, "no-color", false, "Disable ANSI colors")

	rootCmd.AddCommand(
		newDiscoverCmd(),
		newSystemCmd(),
		newZoneCmd(),
		newTunerCmd(),
		newNetusbCmd(),
		newCdCmd(),
		newClockCmd(),
		newDistCmd(),
		newRawCmd(),
		newVersionCmd(),
	)
}

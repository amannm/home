package app

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"github.com/spf13/cobra"
)

func (a *App) Raw(cmd *cobra.Command, args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("raw: requires method and path")
	}
	method := strings.ToUpper(args[0])
	path := args[1]
	queryPairs, err := cmd.Flags().GetStringArray("query")
	if err != nil {
		return err
	}
	data, err := cmd.Flags().GetString("data")
	if err != nil {
		return err
	}
	useStdin, err := cmd.Flags().GetBool("stdin")
	if err != nil {
		return err
	}
	q := url.Values{}
	for _, pair := range queryPairs {
		k, v, ok := splitQuery(pair)
		if ok {
			q.Add(k, v)
		}
	}
	var body []byte
	if useStdin {
		body, err = readStdin()
		if err != nil {
			return err
		}
	} else if strings.TrimSpace(data) != "" {
		if json.Valid([]byte(data)) {
			body = []byte(data)
		} else {
			enc, err := json.Marshal(data)
			if err != nil {
				return err
			}
			body = enc
		}
	}
	return a.call(method, path, q, body, "application/json")
}

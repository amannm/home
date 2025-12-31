package app

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strings"

	"gopkg.in/yaml.v3"
)

func (a *App) render(body []byte) error {
	if len(body) == 0 {
		return nil
	}
	format := strings.ToLower(strings.TrimSpace(a.Options.Format))
	var v any
	if err := json.Unmarshal(body, &v); err != nil {
		_, werr := os.Stdout.Write(body)
		if werr == nil && !bytes.HasSuffix(body, []byte("\n")) {
			_, _ = fmt.Fprintln(os.Stdout)
		}
		return werr
	}
	switch format {
	case "json":
		out, err := json.Marshal(v)
		if err != nil {
			return err
		}
		_, err = os.Stdout.Write(append(out, '\n'))
		return err
	case "yaml":
		out, err := yaml.Marshal(v)
		if err != nil {
			return err
		}
		_, err = os.Stdout.Write(out)
		return err
	case "table":
		out := renderTable(v)
		_, err := os.Stdout.WriteString(out)
		return err
	default:
		out, err := json.MarshalIndent(v, "", "  ")
		if err != nil {
			return err
		}
		_, err = os.Stdout.Write(append(out, '\n'))
		return err
	}
}

func renderTable(v any) string {
	var b strings.Builder
	switch t := v.(type) {
	case map[string]any:
		keys := make([]string, 0, len(t))
		for k := range t {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		for _, k := range keys {
			b.WriteString(k)
			b.WriteString("\t")
			b.WriteString(valueString(t[k]))
			b.WriteString("\n")
		}
	case []any:
		if len(t) == 0 {
			return ""
		}
		allMaps := true
		keySet := map[string]struct{}{}
		for _, item := range t {
			m, ok := item.(map[string]any)
			if !ok {
				allMaps = false
				break
			}
			for k := range m {
				keySet[k] = struct{}{}
			}
		}
		if allMaps {
			keys := make([]string, 0, len(keySet))
			for k := range keySet {
				keys = append(keys, k)
			}
			sort.Strings(keys)
			b.WriteString(strings.Join(keys, "\t"))
			b.WriteString("\n")
			for _, item := range t {
				m := item.(map[string]any)
				for i, k := range keys {
					if i > 0 {
						b.WriteString("\t")
					}
					b.WriteString(valueString(m[k]))
				}
				b.WriteString("\n")
			}
			return b.String()
		}
		for _, item := range t {
			b.WriteString(valueString(item))
			b.WriteString("\n")
		}
	default:
		b.WriteString(valueString(t))
		b.WriteString("\n")
	}
	return b.String()
}

func valueString(v any) string {
	switch t := v.(type) {
	case nil:
		return ""
	case string:
		return t
	case float64:
		if t == float64(int64(t)) {
			return fmt.Sprintf("%d", int64(t))
		}
		return fmt.Sprintf("%v", t)
	case bool:
		if t {
			return "true"
		}
		return "false"
	default:
		out, err := json.Marshal(t)
		if err != nil {
			return fmt.Sprintf("%v", t)
		}
		return string(out)
	}
}

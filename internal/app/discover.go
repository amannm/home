package app

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"net"
	"os/exec"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

type discoveredDevice struct {
	Name      string   `json:"name"`
	Type      string   `json:"type"`
	Domain    string   `json:"domain"`
	Host      string   `json:"host"`
	Port      int      `json:"port"`
	Addresses []string `json:"addresses,omitempty"`
	BaseURL   string   `json:"base_url"`
}

func (a *App) Discover(cmd *cobra.Command, args []string) error {
	devs, err := browseMusicCast()
	if err != nil {
		return err
	}
	out, err := json.Marshal(devs)
	if err != nil {
		return err
	}
	return a.render(out, "application/json")
}

func browseMusicCast() ([]discoveredDevice, error) {
	types := []string{"_musiccast._tcp", "_yamaha._tcp", "_yxc._tcp"}
	all := []discoveredDevice{}
	seen := map[string]struct{}{}
	var lastErr error
	for _, t := range types {
		items, err := browseService(t, 3*time.Second)
		if err != nil {
			lastErr = err
			continue
		}
		for _, item := range items {
			key := item.Name + "|" + item.Type + "|" + item.Domain
			if _, ok := seen[key]; ok {
				continue
			}
			seen[key] = struct{}{}
			all = append(all, item)
		}
	}
	sort.Slice(all, func(i, j int) bool {
		if all[i].Name == all[j].Name {
			return all[i].Host < all[j].Host
		}
		return all[i].Name < all[j].Name
	})
	if len(all) == 0 && lastErr != nil {
		return nil, lastErr
	}
	return all, nil
}

func browseService(service string, timeout time.Duration) ([]discoveredDevice, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	cmd := exec.CommandContext(ctx, "dns-sd", "-B", service)
	out, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}
	if err := cmd.Start(); err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(out)
	type entry struct {
		Name   string
		Type   string
		Domain string
	}
	entries := []entry{}
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "Browsing") || strings.HasPrefix(line, "DATE:") {
			continue
		}
		fields := strings.Fields(line)
		if len(fields) < 6 {
			continue
		}
		action := fields[1]
		if action != "Add" {
			continue
		}
		domain := fields[4]
		svc := fields[5]
		name := strings.Join(fields[6:], " ")
		if name == "" {
			continue
		}
		entries = append(entries, entry{Name: name, Type: svc, Domain: domain})
	}
	_ = cmd.Process.Kill()
	_ = cmd.Wait()
	if len(entries) == 0 {
		return []discoveredDevice{}, nil
	}
	results := []discoveredDevice{}
	for _, e := range entries {
		dev, err := resolveService(e.Name, e.Type, e.Domain, 2*time.Second)
		if err == nil {
			results = append(results, dev)
		}
	}
	return results, nil
}

func resolveService(name, service, domain string, timeout time.Duration) (discoveredDevice, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	cmd := exec.CommandContext(ctx, "dns-sd", "-L", name, service, domain)
	out, err := cmd.Output()
	if err != nil {
		return discoveredDevice{}, err
	}
	host, port := parseResolveOutput(out)
	host = strings.TrimSuffix(host, ".")
	addrs := []string{}
	if host != "" {
		if ips, err := net.LookupIP(host); err == nil {
			for _, ip := range ips {
				addrs = append(addrs, ip.String())
			}
		}
	}
	baseURL := ""
	if host != "" {
		if port == 0 || port == 80 {
			baseURL = "http://" + host + "/YamahaExtendedControl"
		} else {
			baseURL = "http://" + host + ":" + strconv.Itoa(port) + "/YamahaExtendedControl"
		}
	}
	return discoveredDevice{
		Name:      name,
		Type:      service,
		Domain:    domain,
		Host:      host,
		Port:      port,
		Addresses: addrs,
		BaseURL:   baseURL,
	}, nil
}

func parseResolveOutput(out []byte) (string, int) {
	scanner := bufio.NewScanner(bytes.NewReader(out))
	for scanner.Scan() {
		line := scanner.Text()
		idx := strings.Index(line, "can be reached at")
		if idx == -1 {
			continue
		}
		rest := strings.TrimSpace(line[idx+len("can be reached at"):])
		fields := strings.Fields(rest)
		if len(fields) == 0 {
			continue
		}
		hostPort := fields[0]
		hostPort = strings.TrimSuffix(hostPort, ".")
		if strings.Contains(hostPort, ":") {
			parts := strings.Split(hostPort, ":")
			host := strings.Join(parts[:len(parts)-1], ":")
			portStr := parts[len(parts)-1]
			if port, err := strconv.Atoi(portStr); err == nil {
				return host, port
			}
			return host, 0
		}
		return hostPort, 0
	}
	return "", 0
}

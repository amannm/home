package app

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

func (a *App) api(path string) string {
	prefix := strings.TrimSpace(a.Options.APIPrefix)
	if prefix == "" {
		prefix = "/v1"
	}
	if !strings.HasPrefix(prefix, "/") {
		prefix = "/" + prefix
	}
	prefix = strings.TrimRight(prefix, "/")
	path = strings.TrimLeft(path, "/")
	return prefix + "/" + path
}

func (a *App) baseURL() (string, error) {
	if strings.TrimSpace(a.Options.BaseURL) != "" {
		return strings.TrimRight(a.Options.BaseURL, "/"), nil
	}
	host := strings.TrimSpace(a.Options.Host)
	if host == "" {
		return "", errors.New("host or base-url is required")
	}
	if strings.HasPrefix(host, "http://") || strings.HasPrefix(host, "https://") {
		return strings.TrimRight(host, "/") + "/YamahaExtendedControl", nil
	}
	return "http://" + host + "/YamahaExtendedControl", nil
}

func (a *App) get(path string, q url.Values) error {
	return a.call(http.MethodGet, path, q, nil, "")
}

func (a *App) post(path string, q url.Values, body []byte) error {
	return a.call(http.MethodPost, path, q, body, "application/json")
}

func (a *App) call(method, path string, q url.Values, body []byte, contentType string) error {
	req, err := a.buildRequest(context.Background(), method, path, q, body, contentType)
	if err != nil {
		return err
	}
	if a.Options.DryRun {
		return a.printRequest(req, body)
	}
	respBody, status, _, err := a.doRequest(method, path, q, body, contentType)
	if err != nil {
		return err
	}
	if a.Options.Verbose > 0 && !a.Options.Quiet {
		_, _ = fmt.Fprintf(os.Stderr, "%s %s -> %d\n", req.Method, req.URL.String(), status)
	}
	if err := a.render(respBody); err != nil {
		return err
	}
	if status >= 400 {
		return fmt.Errorf("http %d", status)
	}
	if code, ok := responseCode(respBody); ok && code != 0 {
		return fmt.Errorf("response_code %d", code)
	}
	return nil
}

func (a *App) buildRequest(ctx context.Context, method, path string, q url.Values, body []byte, contentType string) (*http.Request, error) {
	u, err := a.buildURL(path, q)
	if err != nil {
		return nil, err
	}
	var reader io.Reader
	if len(body) > 0 {
		reader = bytes.NewReader(body)
	}
	req, err := http.NewRequestWithContext(ctx, method, u, reader)
	if err != nil {
		return nil, err
	}
	if contentType != "" && len(body) > 0 {
		req.Header.Set("Content-Type", contentType)
	}
	for _, h := range a.Options.Headers {
		k, v, ok := splitHeader(h)
		if ok {
			req.Header.Add(k, v)
		}
	}
	if strings.TrimSpace(a.Options.Auth) != "" && req.Header.Get("Authorization") == "" {
		user, pass := splitAuth(a.Options.Auth)
		req.SetBasicAuth(user, pass)
	}
	return req, nil
}

func (a *App) buildURL(path string, q url.Values) (string, error) {
	base := ""
	if strings.HasPrefix(path, "http://") || strings.HasPrefix(path, "https://") {
		base = path
	} else {
		root, err := a.baseURL()
		if err != nil {
			return "", err
		}
		if !strings.HasPrefix(path, "/") {
			path = "/" + path
		}
		base = root + path
	}
	u, err := url.Parse(base)
	if err != nil {
		return "", err
	}
	if q != nil && len(q) > 0 {
		existing := u.Query()
		for k, vals := range q {
			for _, v := range vals {
				existing.Add(k, v)
			}
		}
		u.RawQuery = existing.Encode()
	}
	return u.String(), nil
}

func (a *App) doRequest(method, path string, q url.Values, body []byte, contentType string) ([]byte, int, http.Header, error) {
	retries := a.Options.Retries
	var lastErr error
	for i := 0; i <= retries; i++ {
		ctx := context.Background()
		var cancel context.CancelFunc
		if a.Options.Timeout > 0 {
			ctx, cancel = context.WithTimeout(ctx, a.Options.Timeout)
		}
		req, err := a.buildRequest(ctx, method, path, q, body, contentType)
		if err != nil {
			if cancel != nil {
				cancel()
			}
			return nil, 0, nil, err
		}
		resp, err := http.DefaultClient.Do(req)
		if cancel != nil {
			cancel()
		}
		if err != nil {
			lastErr = err
		} else {
			respBody, rerr := io.ReadAll(resp.Body)
			_ = resp.Body.Close()
			if rerr != nil {
				return nil, resp.StatusCode, resp.Header, rerr
			}
			if resp.StatusCode >= 500 && i < retries {
				time.Sleep(time.Duration(200*(i+1)) * time.Millisecond)
				continue
			}
			return respBody, resp.StatusCode, resp.Header, nil
		}
		if i < retries {
			time.Sleep(time.Duration(200*(i+1)) * time.Millisecond)
		}
	}
	if lastErr != nil {
		return nil, 0, nil, lastErr
	}
	return nil, 0, nil, errors.New("request failed")
}

func (a *App) printRequest(req *http.Request, body []byte) error {
	var b strings.Builder
	b.WriteString(req.Method)
	b.WriteString(" ")
	b.WriteString(req.URL.String())
	b.WriteString("\n")
	for k, vals := range req.Header {
		for _, v := range vals {
			b.WriteString(k)
			b.WriteString(": ")
			b.WriteString(v)
			b.WriteString("\n")
		}
	}
	if len(body) > 0 {
		b.WriteString("\n")
		b.Write(body)
		b.WriteString("\n")
	}
	_, err := os.Stdout.WriteString(b.String())
	return err
}

func responseCode(body []byte) (int, bool) {
	if len(body) == 0 {
		return 0, false
	}
	var v map[string]any
	if err := json.Unmarshal(body, &v); err != nil {
		return 0, false
	}
	val, ok := v["response_code"]
	if !ok {
		return 0, false
	}
	switch t := val.(type) {
	case float64:
		return int(t), true
	case int:
		return t, true
	case string:
		if n, err := strconv.Atoi(t); err == nil {
			return n, true
		}
	}
	return 0, false
}

func splitHeader(h string) (string, string, bool) {
	if strings.Contains(h, ":") {
		parts := strings.SplitN(h, ":", 2)
		return strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1]), parts[0] != ""
	}
	if strings.Contains(h, "=") {
		parts := strings.SplitN(h, "=", 2)
		return strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1]), parts[0] != ""
	}
	return "", "", false
}

func splitAuth(auth string) (string, string) {
	parts := strings.SplitN(auth, ":", 2)
	if len(parts) == 1 {
		return parts[0], ""
	}
	return parts[0], parts[1]
}

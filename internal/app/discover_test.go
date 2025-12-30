package app

import (
	"os"
	"testing"
)

func TestDiscoverLive(t *testing.T) {
	if os.Getenv("CI") != "" {
		t.Skip("set CI=1 to skip")
	}
	devs, err := browseMusicCast()
	if err != nil {
		t.Fatalf("discover failed: %v", err)
	}
	if len(devs) == 0 {
		t.Fatalf("no devices discovered")
	}
	for _, d := range devs {
		if d.Name == "" {
			t.Fatalf("device name empty")
		}
		if d.Host == "" {
			t.Fatalf("host empty for %s", d.Name)
		}
		if d.BaseURL == "" {
			t.Fatalf("base url empty for %s", d.Name)
		}
		t.Logf("%s %s %s:%d %s", d.Name, d.Type, d.Host, d.Port, d.BaseURL)
	}
}

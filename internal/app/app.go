package app

import "time"

type Options struct {
	Host      string
	BaseURL   string
	APIPrefix string
	Zone      string
	Timeout   time.Duration
	Retries   int
	Auth      string
	Headers   []string
	DryRun    bool
	Format    string
	Verbose   int
	Quiet     bool
	NoColor   bool
}

type App struct {
	Options Options
}

func New(opts Options) *App {
	return &App{Options: opts}
}

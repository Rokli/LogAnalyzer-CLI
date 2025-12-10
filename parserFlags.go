package main

import (
	"flag"
	"fmt"
	"os"
)

type Config struct {
	File   string
	Stats  bool
	Level  string
	Search string
	Output string
	Limit  int
	Help   bool
}

func validateConfig(cfg Config) error {
	if cfg.Help {
		return nil
	}

	if cfg.File == "" {
		return fmt.Errorf("file is required")
	}

	if _, err := os.Stat(cfg.File); err != nil {
		return fmt.Errorf("file not found: %w", err)
	}

	return nil
}

func parseFlags() Config {
	var cfg Config
	flag.StringVar(&cfg.File, "file", "", "path to log file")
	flag.BoolVar(&cfg.Stats, "stats", false, "show statistics")
	flag.StringVar(&cfg.Level, "level", "", "level info")
	flag.StringVar(&cfg.Search, "search", "", "find substr")
	flag.StringVar(&cfg.Output, "output", "", "create file json/csv")
	flag.IntVar(&cfg.Limit, "limit", 0, "output with limit")
	flag.BoolVar(&cfg.Help, "help", false, "info about cli")

	flag.Parse()

	return cfg
}

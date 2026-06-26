package ui

import (
	"github.com/shohei81/terminal-setup/internal/config"
)

type moduleMeta struct {
	Key     string
	Desc    string
	Default bool
}

var fastfetchModules = []moduleMeta{
	{"OS",      "macOS version (e.g. macOS Tahoe 26.5)", true},
	{"Host",    "Your Mac model (e.g. MacBook Pro 14-inch, M4)", true},
	{"Shell",   "Which shell you're using (zsh)", true},
	{"CPU",     "Processor name and clock speed", true},
	{"Memory",  "RAM in use vs total (e.g. 12 GB / 16 GB)", true},
	{"Battery", "Battery level and estimated time remaining", true},
	{"Disk",    "Available storage space", false},
	{"GPU",     "Graphics chip name", false},
	{"Uptime",  "How long your Mac has been running", false},
	{"Colors",  "Color palette swatch", true},
}

func RunFastfetch(cfg *config.Config) error {
	isFirstRun := len(cfg.FastfetchModules) == 0
	prev := make(map[string]bool)
	for _, m := range cfg.FastfetchModules {
		prev[m] = true
	}

	items := make([]ChecklistItem, len(fastfetchModules))
	for i, m := range fastfetchModules {
		sel := m.Default
		if !isFirstRun {
			sel = prev[m.Key]
		}
		items[i] = ChecklistItem{Key: m.Key, Desc: m.Desc, Selected: sel}
	}

	result, err := RunChecklist(
		"Terminal Info Display",
		"Choose what to show when a new terminal opens",
		items,
	)
	if err != nil {
		return err
	}

	selected := []string{}
	for _, item := range result {
		if item.Selected {
			selected = append(selected, item.Key)
		}
	}
	cfg.FastfetchModules = selected
	return nil
}

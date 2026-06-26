package ui

import (
	"github.com/charmbracelet/huh"
	"github.com/shohei81/terminal-setup/internal/config"
)

var allModules = []string{
	"OS", "Host", "Shell", "Terminal", "CPU", "GPU",
	"Memory", "Disk", "Battery", "Uptime", "Wifi", "TerminalFont", "Colors",
}

func RunFastfetch(cfg *config.Config) error {
	selected := cfg.FastfetchModules

	err := huh.NewForm(
		huh.NewGroup(
			huh.NewMultiSelect[string]().
				Title("Terminal Info Display").
				Description("Choose what to show when a new terminal window opens").
				Options(moduleOptions(allModules, selected)...).
				Value(&selected),
		),
	).Run()
	if err != nil {
		return err
	}

	cfg.FastfetchModules = selected
	return nil
}

func moduleOptions(all, selected []string) []huh.Option[string] {
	opts := make([]huh.Option[string], len(all))
	for i, m := range all {
		opt := huh.NewOption(m, m)
		for _, s := range selected {
			if s == m {
				opt = opt.Selected(true)
				break
			}
		}
		opts[i] = opt
	}
	return opts
}

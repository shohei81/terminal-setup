package ui

import (
	"fmt"

	"github.com/charmbracelet/huh"
	"github.com/shohei81/terminal-setup/internal/config"
	"github.com/shohei81/terminal-setup/internal/installer"
)

func RunTools(cfg *config.Config) error {
	selected := cfg.OptionalTools

	opts := make([]huh.Option[string], len(installer.TierBMeta))
	for i, t := range installer.TierBMeta {
		label := fmt.Sprintf("%-10s — %s", t.Key, t.Desc)
		opt := huh.NewOption(label, t.Key)
		for _, s := range selected {
			if s == t.Key {
				opt = opt.Selected(true)
				break
			}
		}
		opts[i] = opt
	}

	err := huh.NewForm(
		huh.NewGroup(
			huh.NewMultiSelect[string]().
				Title("Optional Tools").
				Description("Select additional tools to install (you can add more later)").
				Options(opts...).
				Value(&selected),
		),
	).Run()
	if err != nil {
		return err
	}

	cfg.OptionalTools = selected
	return nil
}

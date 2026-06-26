package ui

import (
	"github.com/shohei81/terminal-setup/internal/config"
	"github.com/shohei81/terminal-setup/internal/installer"
)

func RunTools(cfg *config.Config) error {
	prev := make(map[string]bool)
	for _, t := range cfg.OptionalTools {
		prev[t] = true
	}

	items := make([]ChecklistItem, len(installer.TierBMeta))
	for i, t := range installer.TierBMeta {
		items[i] = ChecklistItem{Key: t.Key, Desc: t.Desc, Selected: prev[t.Key]}
	}

	result, err := RunChecklist(
		"Optional Tools",
		"Choose additional tools to install",
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
	cfg.OptionalTools = selected
	return nil
}

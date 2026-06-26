package ui

import (
	"os"
	"path/filepath"

	"github.com/charmbracelet/huh"
	"github.com/shohei81/terminal-setup/internal/config"
)

func RunZshrcStrategy(cfg *config.Config) error {
	home := config.HomeDir()
	zshrcPath := filepath.Join(home, ".zshrc")

	if _, err := os.Stat(zshrcPath); os.IsNotExist(err) {
		cfg.ZshrcStrategy = "overwrite"
		return nil
	}

	strategy := cfg.ZshrcStrategy

	err := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Existing .zshrc detected").
				Description("How should we handle your existing shell configuration?").
				Options(
					huh.NewOption("Append  — add terminal-setup block (recommended, keeps your config)", "append"),
					huh.NewOption("Backup and overwrite  — save backup then replace file", "backup_overwrite"),
					huh.NewOption("Overwrite  — replace file without backup", "overwrite"),
					huh.NewOption("Skip  — leave .zshrc unchanged", "skip"),
				).
				Value(&strategy),
		),
	).Run()
	if err != nil {
		return err
	}

	cfg.ZshrcStrategy = strategy
	return nil
}

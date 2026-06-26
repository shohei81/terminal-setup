package installer

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/shohei81/terminal-setup/internal/config"
	"github.com/shohei81/terminal-setup/internal/templates"
)

var TierA = []struct {
	Pkg  string
	Cask bool
}{
	{"zsh-autosuggestions", false},
	{"zsh-completions", false},
	{"starship", false},
	{"lsd", false},
	{"bat", false},
	{"fd", false},
	{"ripgrep", false},
	{"fastfetch", false},
}

var TierBMeta = []struct {
	Key  string
	Pkg  string
	Desc string
}{
	{"fzf", "fzf", "Quickly search command history and files"},
	{"zoxide", "zoxide", "Jump to frequently used folders instantly"},
	{"yazi", "yazi", "Visual file manager inside the terminal"},
	{"lazygit", "lazygit", "Manage Git visually in the terminal"},
	{"bottom", "bottom", "Beautiful system monitor (CPU, memory, processes)"},
}

func InstallTierA() error {
	fmt.Println("Installing core tools...")
	for _, t := range TierA {
		if err := EnsureInstalled(t.Pkg, t.Cask); err != nil {
			return err
		}
	}
	return installStarshipConfig()
}

func InstallTierB(tools []string) error {
	if len(tools) == 0 {
		return nil
	}
	fmt.Println("Installing optional tools...")
	for _, key := range tools {
		for _, meta := range TierBMeta {
			if meta.Key == key {
				if err := EnsureInstalled(meta.Pkg, false); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func installStarshipConfig() error {
	home := config.HomeDir()
	dir := filepath.Join(home, ".config")
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}
	dst := filepath.Join(dir, "starship.toml")
	if err := os.WriteFile(dst, []byte(templates.StarshipToml), 0644); err != nil {
		return err
	}
	fmt.Println("  ✓ starship.toml written")
	return nil
}

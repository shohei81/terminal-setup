package ui

import (
	"fmt"

	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
	"github.com/shohei81/terminal-setup/internal/config"
)

type ThemeMeta struct {
	Name   string
	Colors [6]string
	Type   string
}

var Themes = []ThemeMeta{
	{"TokyoNight", [6]string{"#1a1b26", "#c0caf5", "#f7768e", "#9ece6a", "#e0af68", "#7aa2f7"}, "dark"},
	{"Catppuccin Mocha", [6]string{"#1e1e2e", "#cdd6f4", "#f38ba8", "#a6e3a1", "#f9e2af", "#89b4fa"}, "dark"},
	{"Dracula", [6]string{"#282a36", "#f8f8f2", "#ff5555", "#50fa7b", "#f1fa8c", "#6272a4"}, "dark"},
	{"Nord", [6]string{"#2e3440", "#d8dee9", "#bf616a", "#a3be8c", "#ebcb8b", "#5e81ac"}, "dark"},
	{"Gruvbox Dark", [6]string{"#282828", "#ebdbb2", "#cc241d", "#98971a", "#d79921", "#458588"}, "dark"},
	{"Rose Pine", [6]string{"#191724", "#e0def4", "#eb6f92", "#31748f", "#f6c177", "#c4a7e7"}, "dark"},
	{"Catppuccin Latte", [6]string{"#eff1f5", "#4c4f69", "#d20f39", "#40a02b", "#df8e1d", "#1e66f5"}, "light"},
	{"Nord Light", [6]string{"#eceff4", "#2e3440", "#bf616a", "#a3be8c", "#ebcb8b", "#5e81ac"}, "light"},
	{"Rose Pine Dawn", [6]string{"#faf4ed", "#575279", "#b4637a", "#286983", "#ea9d34", "#56949f"}, "light"},
	{"GitHub Light Default", [6]string{"#ffffff", "#24292f", "#cf222e", "#1a7f37", "#9a6700", "#0969da"}, "light"},
}

func swatch(colors [6]string) string {
	result := ""
	for _, c := range colors {
		result += lipgloss.NewStyle().Background(lipgloss.Color(c)).Render("  ")
	}
	return result
}

func themeOptions(current string) []huh.Option[string] {
	opts := make([]huh.Option[string], len(Themes))
	for i, t := range Themes {
		label := fmt.Sprintf("%-22s %s  %s", t.Name, swatch(t.Colors), t.Type)
		opts[i] = huh.NewOption(label, t.Name)
		if t.Name == current {
			opts[i] = opts[i].Selected(true)
		}
	}
	return opts
}

func RunAppearance(cfg *config.Config) error {
	theme := cfg.Theme
	fontSizeStr := fmt.Sprintf("%d", cfg.FontSize)
	opacityStr := "solid"
	if cfg.Opacity < 1.0 {
		opacityStr = "transparent"
	}

	err := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Theme").
				Description("Choose a color theme for Ghostty").
				Options(themeOptions(cfg.Theme)...).
				Value(&theme),

			huh.NewSelect[string]().
				Title("Font Size").
				Options(
					huh.NewOption("Small  (12pt)", "12"),
					huh.NewOption("Medium (14pt) — default", "14"),
					huh.NewOption("Large  (16pt)", "16"),
					huh.NewOption("XLarge (18pt)", "18"),
				).
				Value(&fontSizeStr),

			huh.NewSelect[string]().
				Title("Background").
				Options(
					huh.NewOption("Solid (fully opaque)", "solid"),
					huh.NewOption("Slightly transparent", "transparent"),
				).
				Value(&opacityStr),
		),
	).Run()
	if err != nil {
		return err
	}

	cfg.Theme = theme

	switch fontSizeStr {
	case "12":
		cfg.FontSize = 12
	case "14":
		cfg.FontSize = 14
	case "16":
		cfg.FontSize = 16
	case "18":
		cfg.FontSize = 18
	}

	if opacityStr == "transparent" {
		cfg.Opacity = 0.95
	} else {
		cfg.Opacity = 1.0
	}

	return nil
}

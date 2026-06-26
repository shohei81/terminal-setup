package ui

import (
	"fmt"
	"strconv"

	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
	"github.com/shohei81/terminal-setup/internal/config"
)

type ThemeMeta struct {
	Name  string
	Bg    string
	Fg    string
	Red   string
	Green string
	Blue  string
	Type  string
	Vibe  string
}

var Themes = []ThemeMeta{
	{"TokyoNight", "#1a1b26", "#c0caf5", "#f7768e", "#9ece6a", "#7aa2f7", "dark", "Cool blue-purple"},
	{"Catppuccin Mocha", "#1e1e2e", "#cdd6f4", "#f38ba8", "#a6e3a1", "#89b4fa", "dark", "Warm pastel"},
	{"Dracula", "#282a36", "#f8f8f2", "#ff5555", "#50fa7b", "#6272a4", "dark", "Purple & pink accents"},
	{"Nord", "#2e3440", "#d8dee9", "#bf616a", "#a3be8c", "#5e81ac", "dark", "Icy minimal"},
	{"Gruvbox Dark", "#282828", "#ebdbb2", "#cc241d", "#98971a", "#458588", "dark", "Warm earthy retro"},
	{"Rose Pine", "#191724", "#e0def4", "#eb6f92", "#31748f", "#c4a7e7", "dark", "Muted & elegant"},
	{"Catppuccin Latte", "#eff1f5", "#4c4f69", "#d20f39", "#40a02b", "#1e66f5", "light", "Warm pastel, light"},
	{"Nord Light", "#eceff4", "#2e3440", "#bf616a", "#a3be8c", "#5e81ac", "light", "Clean minimal, light"},
	{"Rose Pine Dawn", "#faf4ed", "#575279", "#b4637a", "#286983", "#56949f", "light", "Elegant, light"},
	{"GitHub Light Default", "#ffffff", "#24292f", "#cf222e", "#1a7f37", "#0969da", "light", "Clean & familiar"},
}

// themePreview renders a mini terminal line using the theme's actual colors.
// Example: " $ echo Hello "  with prompt in green, command in blue, text in fg, all on bg.
func themePreview(t ThemeMeta) string {
	base := lipgloss.NewStyle().Background(lipgloss.Color(t.Bg))
	prompt := base.Foreground(lipgloss.Color(t.Green)).Render(" $ ")
	cmd := base.Foreground(lipgloss.Color(t.Blue)).Render("echo")
	text := base.Foreground(lipgloss.Color(t.Fg)).Render(" Hello ")
	pad := base.Render("  ")
	return prompt + cmd + text + pad
}

func themeOptions(current string) []huh.Option[string] {
	opts := make([]huh.Option[string], len(Themes))
	for i, t := range Themes {
		preview := themePreview(t)
		// Fixed-width name column (longest name is "GitHub Light Default" = 20 chars)
		// so 22 gives consistent alignment regardless of name length.
		label := fmt.Sprintf("%-22s  %s  %s", t.Name, preview, t.Vibe)
		opt := huh.NewOption(label, t.Name)
		if t.Name == current {
			opt = opt.Selected(true)
		}
		opts[i] = opt
	}
	return opts
}

func RunAppearance(cfg *config.Config) error {
	theme := cfg.Theme
	fontSizeStr := strconv.Itoa(cfg.FontSize)
	opacityStr := "solid"
	if cfg.Opacity < 1.0 {
		opacityStr = "transparent"
	}

	// Each Group is a separate page in huh — user navigates with Enter / Esc.
	err := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Color Theme").
				Description("Choose a color theme for your terminal").
				Options(themeOptions(theme)...).
				Value(&theme),
		),
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Font Size").
				Description("Choose a comfortable reading size").
				Options(
					huh.NewOption("Small  — 12pt", "12"),
					huh.NewOption("Medium — 14pt (default)", "14"),
					huh.NewOption("Large  — 16pt", "16"),
					huh.NewOption("XLarge — 18pt", "18"),
				).
				Value(&fontSizeStr),
		),
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Background").
				Description("Should the terminal window be slightly see-through?").
				Options(
					huh.NewOption("Solid — fully opaque", "solid"),
					huh.NewOption("Slightly transparent — desktop shows through", "transparent"),
				).
				Value(&opacityStr),
		),
	).WithTheme(HuhTheme()).WithKeyMap(HuhKeyMap()).Run()
	if err != nil {
		return err
	}

	cfg.Theme = theme
	if n, err := strconv.Atoi(fontSizeStr); err == nil {
		cfg.FontSize = n
	}
	if opacityStr == "transparent" {
		cfg.Opacity = 0.95
	} else {
		cfg.Opacity = 1.0
	}
	return nil
}

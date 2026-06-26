package ui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
	"github.com/shohei81/terminal-setup/internal/config"
)

var changeStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#f7768e"))
var sameStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#565f89"))
var addStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#9ece6a"))

func RunConfirm(prev, next config.Config) (bool, error) {
	fmt.Println()
	fmt.Println(headerStyle.Render("  Review & Confirm"))
	fmt.Println()

	printDiff("Theme", prev.Theme, next.Theme)
	printDiff("Font size", fmt.Sprintf("%d", prev.FontSize), fmt.Sprintf("%d", next.FontSize))
	printDiff("Background opacity", fmt.Sprintf("%.2f", prev.Opacity), fmt.Sprintf("%.2f", next.Opacity))
	printDiff("Fastfetch modules",
		strings.Join(prev.FastfetchModules, ", "),
		strings.Join(next.FastfetchModules, ", "))
	printDiff("Optional tools",
		strings.Join(prev.OptionalTools, ", "),
		strings.Join(next.OptionalTools, ", "))
	printDiff(".zshrc strategy", prev.ZshrcStrategy, next.ZshrcStrategy)

	fmt.Println()
	fmt.Println(dimStyle.Render("  Tier A (always installed): Ghostty, Hack Nerd Font, Starship, lsd, bat, fd, ripgrep, fastfetch, zsh plugins"))
	fmt.Println()

	var proceed bool
	err := huh.NewForm(
		huh.NewGroup(
			huh.NewConfirm().
				Title("Install now?").
				Affirmative("Install").
				Negative("Cancel").
				Value(&proceed),
		),
	).Run()
	if err != nil {
		return false, err
	}
	return proceed, nil
}

func printDiff(label, prev, next string) {
	if prev == "" {
		prev = "(none)"
	}
	if next == "" {
		next = "(none)"
	}
	labelPad := fmt.Sprintf("  %-22s", label)
	if prev == next {
		fmt.Println(labelPad + sameStyle.Render(next))
	} else if prev == "(none)" {
		fmt.Println(labelPad + addStyle.Render(next))
	} else {
		fmt.Println(labelPad + changeStyle.Render(prev+" → "+next))
	}
}

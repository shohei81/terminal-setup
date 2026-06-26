package ui

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
)

// HuhTheme returns a custom huh theme using Tokyo Night colors,
// replacing huh's default pink/magenta accents with blue.
func HuhTheme() *huh.Theme {
	t := huh.ThemeBase()

	blue := lipgloss.Color("#7aa2f7")
	subtle := lipgloss.Color("#3b4261")
	fg := lipgloss.Color("#c0caf5")
	muted := lipgloss.Color("#565f89")
	green := lipgloss.Color("#9ece6a")

	t.Focused.Base = t.Focused.Base.BorderForeground(subtle)
	t.Focused.Title = t.Focused.Title.Foreground(blue).Bold(true)
	t.Focused.Description = t.Focused.Description.Foreground(muted)
	t.Focused.SelectSelector = t.Focused.SelectSelector.Foreground(blue)
	t.Focused.Option = t.Focused.Option.Foreground(fg)
	t.Focused.SelectedOption = t.Focused.SelectedOption.Foreground(blue)
	t.Focused.UnselectedOption = t.Focused.UnselectedOption.Foreground(muted)
	t.Focused.SelectedPrefix = t.Focused.SelectedPrefix.Foreground(green)
	t.Focused.UnselectedPrefix = t.Focused.UnselectedPrefix.Foreground(muted)
	t.Focused.MultiSelectSelector = t.Focused.MultiSelectSelector.Foreground(blue)
	t.Focused.TextInput.Cursor = t.Focused.TextInput.Cursor.Foreground(blue)
	t.Focused.Next = t.Focused.Next.Foreground(blue)

	t.Blurred.Title = t.Blurred.Title.Foreground(muted)
	t.Blurred.SelectSelector = t.Blurred.SelectSelector.Foreground(muted)
	t.Blurred.Option = t.Blurred.Option.Foreground(muted)

	return t
}

// HuhKeyMap returns a keymap that adds esc as a quit key alongside ctrl+c.
func HuhKeyMap() *huh.KeyMap {
	km := huh.NewDefaultKeyMap()
	km.Quit = key.NewBinding(key.WithKeys("ctrl+c", "esc"))
	return km
}

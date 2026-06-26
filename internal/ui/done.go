package ui

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

var successStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#9ece6a")).Bold(true)

func ShowDone() {
	fmt.Println()
	fmt.Println(successStyle.Render("  ✓ All done! Your terminal environment is ready."))
	fmt.Println()
	fmt.Println(dimStyle.Render("  Next steps:"))
	fmt.Println(dimStyle.Render("  1. Restart Ghostty to apply all changes"))
	fmt.Println(dimStyle.Render("  2. Open a new terminal window — your new setup will be active"))
	fmt.Println(dimStyle.Render("  3. Run terminal-setup again anytime to change settings"))
	fmt.Println()
}

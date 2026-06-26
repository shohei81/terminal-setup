package ui

import (
	"fmt"
	"os"

	"github.com/charmbracelet/lipgloss"
	"github.com/shohei81/terminal-setup/internal/config"
)

var headerStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#7aa2f7")).
	Bold(true)

var dimStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#565f89"))

var warnStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#e0af68")).
	Bold(true)

func ShowWelcome() {
	fmt.Println()
	fmt.Println(headerStyle.Render(`
  ████████╗███████╗██████╗ ███╗   ███╗██╗███╗   ██╗ █████╗ ██╗
  ╚══██╔══╝██╔════╝██╔══██╗████╗ ████║██║████╗  ██║██╔══██╗██║
     ██║   █████╗  ██████╔╝██╔████╔██║██║██╔██╗ ██║███████║██║
     ██║   ██╔══╝  ██╔══██╗██║╚██╔╝██║██║██║╚██╗██║██╔══██║██║
     ██║   ███████╗██║  ██║██║ ╚═╝ ██║██║██║ ╚████║██║  ██║███████╗
     ╚═╝   ╚══════╝╚═╝  ╚═╝╚═╝     ╚═╝╚═╝╚═╝  ╚═══╝╚═╝  ╚═╝╚══════╝
                              S E T U P`))
	fmt.Println()
	fmt.Println(dimStyle.Render("  Set up a modern terminal environment in minutes."))
	fmt.Println(dimStyle.Render("  This will install: Ghostty, Hack Nerd Font, Starship, and more."))
	fmt.Println()

	if config.IsSandbox() {
		sandbox := os.Getenv("TERMINAL_SETUP_HOME")
		fmt.Println(warnStyle.Render("  ⚠  Sandbox mode: files will be written to " + sandbox))
		fmt.Println()
	}
}

package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/huh"
	"github.com/shohei81/terminal-setup/internal/config"
	"github.com/shohei81/terminal-setup/internal/installer"
	"github.com/shohei81/terminal-setup/internal/ui"
)

func main() {
	// Load previous config (used as defaults)
	prev, err := config.Load()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Warning: could not load previous config: %v\n", err)
		prev = config.DefaultConfig()
	}

	// Working copy for this run
	cfg := prev

	// Welcome
	ui.ShowWelcome()

	var startNow bool
	err = huh.NewForm(
		huh.NewGroup(
			huh.NewConfirm().
				Title("Ready to get started?").
				Affirmative("Let's go").
				Negative("Quit").
				Value(&startNow),
		),
	).WithTheme(ui.HuhTheme()).Run()
	if err != nil || !startNow {
		fmt.Println("Bye!")
		return
	}

	// Appearance
	if err := ui.RunAppearance(&cfg); err != nil {
		exitOnCancel(err)
	}

	// Fastfetch modules
	if err := ui.RunFastfetch(&cfg); err != nil {
		exitOnCancel(err)
	}

	// Optional tools
	if err := ui.RunTools(&cfg); err != nil {
		exitOnCancel(err)
	}

	// .zshrc strategy (only if .zshrc exists)
	if err := ui.RunZshrcStrategy(&cfg); err != nil {
		exitOnCancel(err)
	}

	// Review & Confirm
	proceed, err := ui.RunConfirm(prev, cfg)
	if err != nil {
		exitOnCancel(err)
	}
	if !proceed {
		fmt.Println("Cancelled. No changes were made.")
		return
	}

	// Save config before install
	if err := config.Save(cfg); err != nil {
		fmt.Fprintf(os.Stderr, "Warning: could not save config: %v\n", err)
	}

	// Install
	fmt.Println()
	if err := installer.InstallGhostty(cfg); err != nil {
		fatal(err)
	}
	if err := installer.InstallTierA(); err != nil {
		fatal(err)
	}
	if err := installer.InstallTierB(cfg.OptionalTools); err != nil {
		fatal(err)
	}
	if err := installer.InstallZshrc(cfg); err != nil {
		fatal(err)
	}

	ui.ShowDone()
}

func exitOnCancel(err error) {
	if err == huh.ErrUserAborted || err == ui.ErrAborted {
		fmt.Println("\nCancelled.")
		os.Exit(0)
	}
	fatal(err)
}

func fatal(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

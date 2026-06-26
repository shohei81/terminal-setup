package config

import "os"

// HomeDir returns the effective home directory.
// If TERMINAL_SETUP_HOME is set, use it as a sandbox for development.
func HomeDir() string {
	if sandbox := os.Getenv("TERMINAL_SETUP_HOME"); sandbox != "" {
		return sandbox
	}
	home, _ := os.UserHomeDir()
	return home
}

func IsSandbox() bool {
	return os.Getenv("TERMINAL_SETUP_HOME") != ""
}

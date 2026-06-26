package installer

import (
	"fmt"
	"os/exec"
	"strings"
)

func IsInstalled(pkg string) bool {
	cmd := exec.Command("brew", "list", pkg)
	return cmd.Run() == nil
}

func IsCaskInstalled(pkg string) bool {
	cmd := exec.Command("brew", "list", "--cask", pkg)
	return cmd.Run() == nil
}

func Install(pkg string) error {
	fmt.Printf("  Installing %s...\n", pkg)
	cmd := exec.Command("brew", "install", pkg)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("brew install %s: %w\n%s", pkg, err, strings.TrimSpace(string(out)))
	}
	return nil
}

func InstallCask(pkg string) error {
	fmt.Printf("  Installing %s (cask)...\n", pkg)
	cmd := exec.Command("brew", "install", "--cask", pkg)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("brew install --cask %s: %w\n%s", pkg, err, strings.TrimSpace(string(out)))
	}
	return nil
}

func EnsureInstalled(pkg string, cask bool) error {
	if cask {
		if IsCaskInstalled(pkg) {
			fmt.Printf("  ✓ %s already installed\n", pkg)
			return nil
		}
		return InstallCask(pkg)
	}
	if IsInstalled(pkg) {
		fmt.Printf("  ✓ %s already installed\n", pkg)
		return nil
	}
	return Install(pkg)
}

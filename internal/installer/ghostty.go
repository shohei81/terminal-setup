package installer

import (
	"bytes"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"time"

	"github.com/shohei81/terminal-setup/internal/config"
	"github.com/shohei81/terminal-setup/internal/templates"
)

type GhosttyData struct {
	FontSize int
	Theme    string
	Opacity  float64
}

func InstallGhostty(cfg config.Config) error {
	if err := EnsureInstalled("ghostty", true); err != nil {
		return err
	}
	if err := EnsureInstalled("font-hack-nerd-font", true); err != nil {
		return err
	}
	return writeGhosttyConfig(cfg)
}

func writeGhosttyConfig(cfg config.Config) error {
	home := config.HomeDir()
	dir := filepath.Join(home, ".config", "ghostty")
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	dst := filepath.Join(dir, "config")
	if _, err := os.Stat(dst); err == nil {
		bak := dst + ".bak." + time.Now().Format("20060102")
		data, _ := os.ReadFile(dst)
		if err := os.WriteFile(bak, data, 0644); err != nil {
			return fmt.Errorf("backup ghostty config: %w", err)
		}
		fmt.Printf("  ✓ Ghostty config backed up to %s\n", bak)
	}

	tmpl, err := template.New("ghostty").Parse(templates.GhosttyTmpl)
	if err != nil {
		return err
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, GhosttyData{
		FontSize: cfg.FontSize,
		Theme:    cfg.Theme,
		Opacity:  cfg.Opacity,
	}); err != nil {
		return err
	}

	if err := os.WriteFile(dst, buf.Bytes(), 0644); err != nil {
		return err
	}
	fmt.Printf("  ✓ Ghostty config written\n")
	return nil
}

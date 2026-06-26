package installer

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
	"time"

	"github.com/shohei81/terminal-setup/internal/config"
	"github.com/shohei81/terminal-setup/internal/templates"
)

const markerBegin = "# BEGIN terminal-setup"
const markerEnd = "# END terminal-setup"

type zshrcData struct {
	FastfetchModules  []string
	FastfetchStructure string
	HasFzf            bool
	HasZoxide         bool
	HasYazi           bool
}

func InstallZshrc(cfg config.Config) error {
	home := config.HomeDir()
	path := filepath.Join(home, ".zshrc")

	block, err := renderZshrcBlock(cfg)
	if err != nil {
		return err
	}

	switch cfg.ZshrcStrategy {
	case "backup_overwrite":
		if err := backupFile(path); err != nil {
			return err
		}
		return os.WriteFile(path, []byte(block), 0644)
	case "overwrite":
		return os.WriteFile(path, []byte(block), 0644)
	case "append":
		return appendWithMarker(path, block)
	case "skip":
		fmt.Println("  Skipping .zshrc (as requested)")
		return nil
	}
	return nil
}

func renderZshrcBlock(cfg config.Config) (string, error) {
	tmpl, err := template.New("zshrc").Parse(templates.ZshrcTmpl)
	if err != nil {
		return "", err
	}

	data := zshrcData{
		FastfetchModules:   cfg.FastfetchModules,
		FastfetchStructure: strings.Join(cfg.FastfetchModules, ":"),
		HasFzf:             config.Contains(cfg.OptionalTools, "fzf"),
		HasZoxide:          config.Contains(cfg.OptionalTools, "zoxide"),
		HasYazi:            config.Contains(cfg.OptionalTools, "yazi"),
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", err
	}
	return buf.String(), nil
}

func appendWithMarker(path, block string) error {
	var existing []byte
	if data, err := os.ReadFile(path); err == nil {
		existing = data
	}

	content := string(existing)

	begin := strings.Index(content, markerBegin)
	end := strings.Index(content, markerEnd)

	var result string
	if begin >= 0 && end >= 0 && end > begin {
		// Replace existing marker block
		endLine := end + len(markerEnd)
		if endLine < len(content) && content[endLine] == '\n' {
			endLine++
		}
		result = content[:begin] + block + content[endLine:]
	} else {
		// Append to end
		if len(content) > 0 && !strings.HasSuffix(content, "\n") {
			content += "\n"
		}
		result = content + "\n" + block
	}

	if err := os.WriteFile(path, []byte(result), 0644); err != nil {
		return err
	}
	fmt.Println("  ✓ .zshrc updated")
	return nil
}

func backupFile(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	ext := filepath.Ext(path)
	base := strings.TrimSuffix(path, ext)
	bak := base + ext + ".bak." + time.Now().Format("20060102")
	if err := os.WriteFile(bak, data, 0644); err != nil {
		return fmt.Errorf("backup %s: %w", path, err)
	}
	fmt.Printf("  ✓ Backed up to %s\n", bak)
	return nil
}

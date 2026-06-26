package config

import (
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Theme            string   `toml:"theme"`
	FontSize         int      `toml:"font_size"`
	Opacity          float64  `toml:"opacity"`
	FastfetchModules []string `toml:"fastfetch_modules"`
	OptionalTools    []string `toml:"optional_tools"`
	ZshrcStrategy    string   `toml:"zshrc_strategy"`
}

func DefaultConfig() Config {
	return Config{
		Theme:            "TokyoNight",
		FontSize:         14,
		Opacity:          0.95,
		FastfetchModules: []string{"OS", "Host", "Shell", "CPU", "Memory", "Battery", "Colors"},
		OptionalTools:    []string{},
		ZshrcStrategy:    "append",
	}
}

func ConfigPath() string {
	return filepath.Join(HomeDir(), ".config", "terminal-setup", "config.toml")
}

func Load() (Config, error) {
	cfg := DefaultConfig()
	path := ConfigPath()
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return cfg, nil
	}
	_, err := toml.DecodeFile(path, &cfg)
	return cfg, err
}

func Save(cfg Config) error {
	path := ConfigPath()
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return err
	}
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	return toml.NewEncoder(f).Encode(cfg)
}

func Contains(slice []string, s string) bool {
	for _, v := range slice {
		if v == s {
			return true
		}
	}
	return false
}

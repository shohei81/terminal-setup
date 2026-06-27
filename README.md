# terminal-setup

A CLI tool that sets up a modern macOS terminal environment through an interactive, choice-based TUI — no config files to write.

```
curl -fsSL https://raw.githubusercontent.com/shohei81/terminal-setup/main/install.sh | sh
```

→ [Full guide and docs](https://shohei81.github.io/terminal-setup/)

---

## What it does

Runs an interactive setup wizard that:

1. Lets you pick a **color theme** (10 options with live previews), font size, and background opacity
2. Selects which **fastfetch** system info modules to show on terminal open
3. Optionally installs extra tools (fzf, zoxide, yazi, lazygit, bottom)
4. Writes a config block to your `.zshrc`

Then installs:

- **Ghostty** — GPU-accelerated terminal emulator
- **Hack Nerd Font** — icon-patched monospace font
- **Starship** — fast, informative shell prompt
- **lsd, bat, fd, ripgrep** — modern replacements for ls, cat, find, grep
- **fastfetch** — system info display
- **zsh-autosuggestions, zsh-completions** — shell productivity plugins

## Requirements

- macOS (Apple Silicon or Intel)
- Internet connection
- Homebrew — installed automatically if missing

## Development

```sh
git clone https://github.com/shohei81/terminal-setup
cd terminal-setup
go build ./...

# Run in sandbox (won't touch your real ~/.zshrc or ~/.config)
TERMINAL_SETUP_HOME=/tmp/ts-sandbox ./terminal-setup
```

### Release

```sh
git tag v0.x.x
git push origin v0.x.x
```

GoReleaser builds macOS arm64/amd64 binaries via GitHub Actions and publishes them to GitHub Releases.

## Project structure

```
main.go                      — entry point, orchestrates TUI flow
internal/
  config/                    — config load/save + sandbox mode
  ui/                        — TUI screens (appearance, fastfetch, tools, etc.)
  installer/                 — Homebrew installs, .zshrc writer
  templates/zshrc.tmpl       — generated .zshrc block
docs/                        — GitHub Pages site
install.sh                   — curl-pipe installer
.goreleaser.yaml             — release config
```

## License

MIT

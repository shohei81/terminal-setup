#!/bin/sh
set -e

REPO="shohei81/terminal-setup"
BIN_NAME="terminal-setup"

# Install Homebrew if not present
if ! command -v brew >/dev/null 2>&1; then
  echo "Installing Homebrew..."
  /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
fi

# Detect architecture
ARCH=$(uname -m)
OS=$(uname -s | tr '[:upper:]' '[:lower:]')

# Download latest release binary
LATEST=$(curl -fsSL "https://api.github.com/repos/$REPO/releases/latest" | grep '"tag_name"' | sed 's/.*"tag_name": "\(.*\)".*/\1/')
URL="https://github.com/$REPO/releases/download/$LATEST/${BIN_NAME}_${OS}_${ARCH}.tar.gz"

TMP=$(mktemp -d)
curl -fsSL "$URL" -o "$TMP/terminal-setup.tar.gz"
tar -xzf "$TMP/terminal-setup.tar.gz" -C "$TMP"
sudo mv "$TMP/$BIN_NAME" /usr/local/bin/
rm -rf "$TMP"

echo "Running terminal-setup..."
terminal-setup

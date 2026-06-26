#!/bin/sh
set -e

REPO="shohei81/terminal-setup"
BIN_NAME="terminal-setup"
INSTALL_DIR="/usr/local/bin"

# macOS only
if [ "$(uname -s)" != "Darwin" ]; then
  echo "Error: This tool is for macOS only." >&2
  exit 1
fi

# Install Homebrew if not present
if ! command -v brew >/dev/null 2>&1; then
  echo "Installing Homebrew..."
  /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
  # Add brew to PATH for Apple Silicon
  if [ -f /opt/homebrew/bin/brew ]; then
    eval "$(/opt/homebrew/bin/brew shellenv)"
  fi
fi

# Detect architecture
ARCH=$(uname -m)
case "$ARCH" in
  arm64)  ARCH_NAME="arm64" ;;
  x86_64) ARCH_NAME="amd64" ;;
  *)
    echo "Error: Unsupported architecture: $ARCH" >&2
    exit 1
    ;;
esac

echo "Fetching latest release..."
LATEST=$(curl -fsSL "https://api.github.com/repos/$REPO/releases/latest" \
  | grep '"tag_name"' \
  | sed 's/.*"tag_name": *"\(.*\)".*/\1/')

if [ -z "$LATEST" ]; then
  echo "Error: Could not determine latest release." >&2
  exit 1
fi

echo "Downloading $BIN_NAME $LATEST (darwin/$ARCH_NAME)..."
ARCHIVE="${BIN_NAME}_darwin_${ARCH_NAME}.tar.gz"
URL="https://github.com/$REPO/releases/download/$LATEST/$ARCHIVE"

TMP=$(mktemp -d)
trap 'rm -rf "$TMP"' EXIT

curl -fsSL "$URL" -o "$TMP/$ARCHIVE"
tar -xzf "$TMP/$ARCHIVE" -C "$TMP"

# Install binary
if [ -w "$INSTALL_DIR" ]; then
  mv "$TMP/$BIN_NAME" "$INSTALL_DIR/$BIN_NAME"
  chmod +x "$INSTALL_DIR/$BIN_NAME"
else
  sudo mv "$TMP/$BIN_NAME" "$INSTALL_DIR/$BIN_NAME"
  sudo chmod +x "$INSTALL_DIR/$BIN_NAME"
fi

echo ""
echo "Installed $BIN_NAME $LATEST to $INSTALL_DIR/$BIN_NAME"
echo ""

# Run setup
exec "$INSTALL_DIR/$BIN_NAME"

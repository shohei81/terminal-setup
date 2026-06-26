class TerminalSetup < Formula
  desc "Interactive modern terminal environment setup for macOS"
  homepage "https://github.com/shohei81/terminal-setup"
  version "0.1.0"

  on_macos do
    if Hardware::CPU.arm?
      url "https://github.com/shohei81/terminal-setup/releases/download/v#{version}/terminal-setup_darwin_arm64.tar.gz"
      sha256 "PLACEHOLDER_ARM64_SHA256"
    else
      url "https://github.com/shohei81/terminal-setup/releases/download/v#{version}/terminal-setup_darwin_amd64.tar.gz"
      sha256 "PLACEHOLDER_AMD64_SHA256"
    end
  end

  def install
    bin.install "terminal-setup"
  end

  test do
    system "#{bin}/terminal-setup", "--version"
  end
end

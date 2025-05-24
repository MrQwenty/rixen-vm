class Rixen < Formula
  desc "Developer-first VM manager for macOS"
  homepage "https://github.com/MrQwenty/rixen-vm"
  url "https://github.com/MrQwenty/rixen-vm/releases/download/v1.0.0/rixen.tar.gz"
  sha256 "61b405035496b16c92c27d7c87cc3bc8773800781112a8704a586c029cd1499d"
  version "1.0.0"

  def install
    bin.install "rx"
  end
end

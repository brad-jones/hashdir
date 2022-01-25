class HashdirCli < Formula
  desc "CLI that generates a hash of an entire directory"
  homepage "https://github.com/brad-jones/hashdir"
  url "https://github.com/brad-jones/hashdir/releases/download/v${VERSION}/hashdir_darwin_amd64.tar.gz"
  version "${VERSION}"
  sha256 "${HASH}"

  def install
    bin.install "hashdir"
  end

  test do
    system "#{bin}/hashdir -v"
  end
end

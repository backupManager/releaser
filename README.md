# GoReleaser [![Build Status](https://travis-ci.org/goreleaser/releaser.svg?branch=master)](https://travis-ci.org/goreleaser/releaser) [![Go Report Card](https://goreportcard.com/badge/github.com/goreleaser/releaser)](https://goreportcard.com/report/github.com/goreleaser/releaser) [![Powered By: GoReleaser](https://img.shields.io/badge/powered%20by-goreleaser-green.svg?style=flat-square)](https://github.com/goreleaser)

Deliver Go binaries as fast and easily as possible.

GoReleaser builds Go binaries for several platforms, creates a github release and then
push a homebrew formulae to a repository. All that wrapped in your favorite CI.

## How it works?

The idea started with a [simple shell script](https://github.com/goreleaser/old-go-releaser),
but it quickly became more complex and I also wanted to publish binaries via
homebrew.

So, the all-new goreleaser was born.

## Usage

Basically, you need to create a `goreleaser.yml` file in the root of your
repository. A minimal config would look like this:

```yaml
repo: user/repo
binary_name: my-binary
```

This will build `main.go` file as `my-binary`, for _Darwin_ and _Linux_,
_x86_64_ and _i386_, packaging the binary, `LICENSE.md` and `README.md`
and publish a new github release in the `user/repo` repository with
the `.tar.gz` files there.

### Homebrew

To push it to a homebrew repo, just add a `brew` section:

```yaml
repo: user/repo
binary_name: my-binary
brew:
  repo: user/homebrew-formulae
  caveats: "Optional caveats to add to the formulae"
```

### Build customization

Just add a `build` section

```yaml
repo: user/repo
binary_name: my-binary
build:
  main: ./cmd/main.go
  oses:
    - darwin
    - freebsd
  arches:
    - amd64
```

> `oses` and `arches` should be in `GOOS`/`GOARCH`-compatible format.

### Add more files

You might also want to change the files that are packaged by adding a `files`
section:

```yaml
repo: user/repo
binary_name: my-binary
files:
  - LICENSE.txt
  - README.md
  - CHANGELOG.md
```

## Wire it with travis-ci

You may want to wire this to auto-deploy your new tags on travis, for examepl:

```yaml
after_success:
  test ! -z "$TRAVIS_TAG" && curl -s https://raw.githubusercontent.com/goreleaser/get/master/latest | bash
```

## How the end result looks like

The release on github looks pretty much like this:

[![image](https://cloud.githubusercontent.com/assets/245435/21578845/09404c8a-cf78-11e6-92d7-165ddc03ca6c.png)
](https://github.com/goreleaser/releaser/releases)

And the [homebrew formulae](https://github.com/goreleaser/homebrew-formulae/blob/master/release.rb) would look like:

```rb
class Release < Formula
  desc "Deliver Go binaries as fast and easily as possible."
  homepage "https://github.com/goreleaser/releaser"
  url "https://github.com/goreleaser/releaser/releases/download/v0.2.0/release_#{%x(uname -s).gsub(/\n/, '')}_#{%x(uname -m).gsub(/\n/, '')}.tar.gz"
  head "https://github.com/goreleaser/releaser.git"
  version "v0.2.0"

  def install
    bin.install "release"
  end
end
```

## Badges

Feel free to use it in your own projects:

```md
[![Powered By: GoReleaser](https://img.shields.io/badge/powered%20by-goreleaser-green.svg?style=flat-square)](https://github.com/goreleaser)
```

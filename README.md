# party-clippy

[![Travis CI](https://img.shields.io/travis/jessfraz/party-clippy.svg?style=for-the-badge)](https://travis-ci.org/jessfraz/party-clippy)
[![GoDoc](https://img.shields.io/badge/godoc-reference-5272B4.svg?style=for-the-badge)](https://godoc.org/github.com/jessfraz/party-clippy)

**Table of Contents**

<!-- toc -->

<!-- tocstop -->

## Installation

#### Binaries

For installation instructions from binaries please visit the [Releases Page](https://github.com/jessfraz/party-clippy/releases).

- **darwin** [386](https://github.com/jessfraz/party-clippy/releases/download/v0.2.3/party-clippy-darwin-386) / [amd64](https://github.com/jessfraz/party-clippy/releases/download/v0.2.3/party-clippy-darwin-amd64)
- **freebsd** [386](https://github.com/jessfraz/party-clippy/releases/download/v0.2.3/party-clippy-freebsd-386) / [amd64](https://github.com/jessfraz/party-clippy/releases/download/v0.2.3/party-clippy-freebsd-amd64)
- **linux** [386](https://github.com/jessfraz/party-clippy/releases/download/v0.2.3/party-clippy-linux-386) / [amd64](https://github.com/jessfraz/party-clippy/releases/download/v0.2.3/party-clippy-linux-amd64) / [arm](https://github.com/jessfraz/party-clippy/releases/download/v0.2.3/party-clippy-linux-arm) / [arm64](https://github.com/jessfraz/party-clippy/releases/download/v0.2.3/party-clippy-linux-arm64)
- **solaris** [amd64](https://github.com/jessfraz/party-clippy/releases/download/v0.2.3/party-clippy-solaris-amd64)
- **windows** [386](https://github.com/jessfraz/party-clippy/releases/download/v0.2.3/party-clippy-windows-386) / [amd64](https://github.com/jessfraz/party-clippy/releases/download/v0.2.3/party-clippy-windows-amd64)

#### Run in Docker

```console
$ docker run -d \
    --name clippy \
    -t \ # need the tty for colors
    -p 8080:8080 \
    r.j3ss.co/party-clippy -d
```

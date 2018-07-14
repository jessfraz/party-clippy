# party-clippy

[![Travis CI](https://travis-ci.org/jessfraz/party-clippy.svg?branch=master)](https://travis-ci.org/jessfraz/party-clippy)

## Installation

#### Binaries

- **darwin** [386](https://github.com/jessfraz/party-clippy/releases/download/v0.2.1/party-clippy-darwin-386) / [amd64](https://github.com/jessfraz/party-clippy/releases/download/v0.2.1/party-clippy-darwin-amd64)
- **freebsd** [386](https://github.com/jessfraz/party-clippy/releases/download/v0.2.1/party-clippy-freebsd-386) / [amd64](https://github.com/jessfraz/party-clippy/releases/download/v0.2.1/party-clippy-freebsd-amd64)
- **linux** [386](https://github.com/jessfraz/party-clippy/releases/download/v0.2.1/party-clippy-linux-386) / [amd64](https://github.com/jessfraz/party-clippy/releases/download/v0.2.1/party-clippy-linux-amd64) / [arm](https://github.com/jessfraz/party-clippy/releases/download/v0.2.1/party-clippy-linux-arm) / [arm64](https://github.com/jessfraz/party-clippy/releases/download/v0.2.1/party-clippy-linux-arm64)
- **solaris** [amd64](https://github.com/jessfraz/party-clippy/releases/download/v0.2.1/party-clippy-solaris-amd64)
- **windows** [386](https://github.com/jessfraz/party-clippy/releases/download/v0.2.1/party-clippy-windows-386) / [amd64](https://github.com/jessfraz/party-clippy/releases/download/v0.2.1/party-clippy-windows-amd64)

#### Run in Docker

```console
$ docker run -d \
    --name clippy \
    -t \ # need the tty for colors
    -p 8080:8080 \
    r.j3ss.co/party-clippy -d
```

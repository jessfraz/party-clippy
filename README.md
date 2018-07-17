# party-clippy

[![Travis CI](https://img.shields.io/travis/jessfraz/party-clippy.svg?style=for-the-badge)](https://travis-ci.org/jessfraz/party-clippy)
[![GoDoc](https://img.shields.io/badge/godoc-reference-5272B4.svg?style=for-the-badge)](https://godoc.org/github.com/jessfraz/party-clippy)

 * [Installation](README.md#installation)
      * [Binaries](README.md#binaries)
      * [Run in Docker](README.md#run-in-docker)<Paste>

## Installation

#### Binaries

For installation instructions from binaries please visit the [Releases Page](https://github.com/jessfraz/party-clippy/releases).

#### Run in Docker

```console
$ docker run -d \
    --name clippy \
    -t \ # need the tty for colors
    -p 8080:8080 \
    r.j3ss.co/party-clippy -d
```

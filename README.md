# hget

[![GitHub Workflow Status](https://img.shields.io/github/actions/workflow/status/clementi/hget/ci.yml)](https://github.com/clementi/hget/actions/workflows/ci.yml)
[![GitHub release (latest by date)](https://img.shields.io/github/v/release/clementi/hget)](https://github.com/clementi/hget/releases)

hget is a command-line, multipart, resumable downloader. It is based on [the project by huydx](https://github.com/huydx/hget).

![](https://i.gyazo.com/641166ab79e196e35d1a0ef3f9befd80.png)

## Install

```
go install github.com/clementi/hget@latest
```

### Binaries

Binaries for Windows, Linux, macOS (arm64 and amd64), FreeBSD, NetBSD amd OpenBSD are available at [Releases](https://github.com/clementi/hget/releases).

## Usage

```
NAME:
   hget - Multipart resumable downloads

USAGE:
   hget [global options] command [command options] [URL]

VERSION:
   2.0.0-beta1

AUTHORS:
   huydx (https://github.com/huydx)
   clementi (https://github.com/clementi)

COMMANDS:
   tasks, t  manage current tasks

GLOBAL OPTIONS:
   --connections value, -n value  number of connections (default: 4)
   --skip-tls, -s                 do not verify certificate for HTTPS (default: false)
   --help, -h                     show help (default: false)
   --version, -v                  print the version (default: false)
```

To interrupt any downloading process, just hit ctrl-c or ctrl-d during the download. hget will safely save your data to `$HOME/.hget` and you will be able to resume later.


![Demo](https://github.com/clementi/hget/blob/main/demo.gif)

<!-- ### Download
![](https://i.gyazo.com/89009c7f02fea8cb4cbf07ee5b75da0a.gif)

### Resume
![](https://i.gyazo.com/caa69808f6377421cb2976f323768dc4.gif)
 -->


# Micronet

A tiny network helper that lists your machine's local IPv4 addresses, with optional verbose logging.

## Overview

Micronet inspects the active network interfaces on your machine and prints out the IPv4 addresses reachable from your local network. It is handy when you need to:

- Discover which address your service is bound to on a dev box or VM.
- Double-check connectivity after changing VPN or Wi-Fi settings.
- Share an address quickly with teammates without digging through OS dialogs.

The tool is organized into three focused components:

1. **Logger** – Centralizes formatted INFO/ERROR output and respects the `-v` flag.
2. **IP Finder** – Walks network interfaces, filters out inactive or loopback adapters, and collects IPv4s.
3. **Main** – Parses flags, wires the pieces together, and presents the results.

## Requirements

- Go 1.22 or newer (any recent stable release works).

## Installation

Clone the repo and optionally install the binary:

```bash
git clone https://github.com/yourname/micronet.git
cd micronet
go install .
```

This drops the `micronet` binary in your `$GOBIN` (or `$GOPATH/bin` if unset).

## Usage

Run the tool from the project root (or anywhere after `go install`):

```bash
# Basic output
go run .

# Verbose mode: includes an INFO line when scanning begins
go run . -v
```

Example output:

```
$ go run . -v
INFO: 2024/10/16 18:20:05 Scanning network interfaces...
192.168.1.42
10.0.0.15
```

If no IPv4 addresses are discovered, Micronet prints:

```
[No IPs found!]
```

## Flags

| Flag | Default | Description                     |
|------|---------|---------------------------------|
| `-v` | `false` | Emit verbose INFO log messages. |

## Project Structure

```
.
├── logger.go      # Logger struct and methods (INFO/ERROR)
├── ip_finder.go   # Interface scanning and IPv4 collection
├── main.go        # CLI entry point and wiring
└── go.mod
```

## Development

Format and build:

```bash
go fmt ./...
go build ./...
```

Optional vetting:

```bash
go vet ./...
```

Happy hacking! Let Micronet keep your local IPs at your fingertips.

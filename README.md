# maclookup

A command-line tool for looking up MAC address vendors using the IEEE OUI database.

[![golangci-lint](https://github.com/fmurodov/maclookup/actions/workflows/lint.yml/badge.svg)](https://github.com/fmurodov/maclookup/actions/workflows/lint.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/fmurodov/maclookup)](https://goreportcard.com/report/github.com/fmurodov/maclookup)
[![codecov](https://codecov.io/gh/fmurodov/maclookup/branch/master/graph/badge.svg?token=S70X7FR02U)](https://codecov.io/gh/fmurodov/maclookup)

## Features

- Fast MAC address vendor lookup
- Automatic OUI database download and caching
- Support for various MAC address formats
- Simple and intuitive command-line interface

## Installation

```bash
go install github.com/fmurodov/maclookup@latest
```

## Usage

```bash
maclookup MAC_ADDRESS
```

### Examples

```bash
# Using colon separator
$ maclookup 74:4d:28:ab:ba:ab
74:4d:28:ab:ba:ab Routerboard.com

# Using hyphen separator
$ maclookup 74-4d-28-ab-ba-ab
74-4d-28-ab-ba-ab Routerboard.com

# Using dot separator
$ maclookup 74.4d.28.ab.ba.ab
74.4d.28.ab.ba.ab Routerboard.com
```

## How it Works

1. The tool first checks if the OUI database is cached locally
2. If not found, it downloads the latest OUI database from IEEE
3. The MAC address is normalized and the vendor is looked up
4. The result is displayed in a simple format

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

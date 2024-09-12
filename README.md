# GODO

> Predefined task runner.

> [!IMPORTANT]
>
> `GODO` is still in the early stages of development.

<br />

[![Go Reference](https://pkg.go.dev/badge/github.com/bastean/godo.svg)](https://pkg.go.dev/github.com/bastean/godo)
[![GitHub Releases](https://img.shields.io/github/v/release/bastean/godo.svg)](https://github.com/bastean/godo/releases)

[![Upgrade workflow](https://github.com/bastean/godo/actions/workflows/upgrade.yml/badge.svg)](https://github.com/bastean/godo/actions/workflows/upgrade.yml)
[![CI workflow](https://github.com/bastean/godo/actions/workflows/ci.yml/badge.svg)](https://github.com/bastean/godo/actions/workflows/ci.yml)
[![Release workflow](https://github.com/bastean/godo/actions/workflows/release.yml/badge.svg)](https://github.com/bastean/godo/actions/workflows/release.yml)

[![License MIT](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/bastean/godo)](https://goreportcard.com/report/github.com/bastean/godo)
[![Commitizen friendly](https://img.shields.io/badge/commitizen-friendly-brightgreen.svg)](https://github.com/commitizen/cz-cli)
[![Release It!](https://img.shields.io/badge/%F0%9F%93%A6%F0%9F%9A%80-release--it-orange.svg)](https://github.com/release-it/release-it)

## CLI

### Installation

```bash
go install github.com/bastean/godo/cmd/godo@latest
```

### Usage

```bash
godo -h
```

```text
________________  ________ _______
__  ____/__  __ \ ___  __ \__  __ \
_  / __  _  / / / __  / / /_  / / /
/ /_/ /  / /_/ /  _  /_/ / / /_/ /
\____/   \____/   /_____/  \____/

Predefined task runner.

Usage: godo [flags]

  -task string
    	Name of the task to run (required)
```

## Package

### Download

```bash
go get github.com/bastean/godo
```

### Update

```bash
go get -u github.com/bastean/godo
```

### Usage

```go
package main

import (
	"github.com/bastean/godo"
)

func main() {
	godo.TODO()
}
```

## Tasks

### Run

-

### Sync

-

### Copy

-

## First Steps

### Clone

#### HTTPS

```bash
git clone https://github.com/bastean/godo.git && cd godo
```

#### SSH

```bash
git clone git@github.com:bastean/godo.git && cd godo
```

### Initialize

#### Locally

1. System Requirements

   - [Go](https://go.dev/doc/install)
   - [Node](https://nodejs.org/en/download)
   - [Make](https://www.gnu.org/software/make)

2. Run

   ```bash
   make init
   ```

### Run

#### Tests

##### Unit

```bash
make test-unit
```

##### Integration

```bash
make test-integration
```

##### Acceptance

```bash
make test-acceptance
```

##### Unit / Integration / Acceptance

```bash
make tests
```

## Tech Stack

#### Base

- [Go](https://go.dev)

#### Please see

- [go.mod](go.mod)
- [package.json](package.json)

## Contributing

- Contributions and Feedback are always welcome!
  - [Open a new issue](https://github.com/bastean/godo/issues/new/choose)

## License

- [MIT](LICENSE)

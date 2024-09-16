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
godo
```

```text
________________  ________ _______
__  ____/__  __ \ ___  __ \__  __ \
_  / __  _  / / / __  / / /_  / / /
/ /_/ /  / /_/ /  _  /_/ / / /_/ /
\____/   \____/   /_____/  \____/

Predefined task runner.

Usage:
  godo [flags]
  godo [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  exec        Execute a list of tasks from a file
  help        Help about any command

Flags:
  -h, --help   help for godo

Use "godo [command] --help" for more information about a command.
```

## Tasks

### Exec

Execute a list of tasks from a file, all tasks will be executed following the preset order in the configuration file, if any of the tasks fail, the following tasks will not be executed.

```bash
godo exec -c configs/example.exec.json
```

```text
Example tasks started!
Create a file
We will create a file and write something in it
(1/2) mkdir -p ignore
(2/2) bash -c echo "Example" > ignore/example.txt
File created!
Example tasks completed!
```

[example.exec.json](configs/example.exec.json)

```json
{
  "Tasks": [
    {
      "Success": "Example tasks started!"
    },
    {
      "Title": "Create a file",
      "Description": "We will create a file and write something in it",
      "Commands": [
        {
          "Name": "mkdir",
          "Args": ["-p", "ignore"]
        },
        {
          "Name": "bash",
          "Args": ["-c", "echo \"Example\" > ignore/example.txt"]
        }
      ],
      "Success": "File created!",
      "Error": "File could not be created"
    },
    {
      "Success": "Example tasks completed!"
    }
  ]
}
```

Basic structure

```json
{
  "Tasks": [
    {
      "Title": "",
      "Description": "",
      "Commands": [{ "Name": "", "Args": [""] }],
      "Success": "",
      "Error": ""
    }
  ]
}
```

- `Success` is only displayed if the task has been successfully completed.
- `Error` is only displayed if an error occurred during the execution of the task.
- All fields in a task are optional.

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

##### Unit & Integration

```bash
make tests
```

## Tech Stack

#### Base

- [Go](https://go.dev)
- [Cobra](https://cobra.dev)

#### Please see

- [go.mod](go.mod)
- [package.json](package.json)

## Contributing

- Contributions and Feedback are always welcome!
  - [Open a new issue](https://github.com/bastean/godo/issues/new/choose)

## License

- [MIT](LICENSE)

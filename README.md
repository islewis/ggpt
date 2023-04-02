# ggpt
<img src="https://github.com/islewis/ggpt/logo/logo.png" width="100">
ggpt is a simple tool for accessing GPT on the command line, written in Go.

## Installation

Make sure go is installed using `go version`. If not, [install go](https://go.dev/doc/install).
```
$ go version      
go version go1.20.0 linux/amd64
```
Next, install ggpt using `go install`:
```
$ go install github.com/islewis/ggpt@latest
```
Thats it! To confirm ggpt is downloaded, run `ggpt --help`:
```
$ ggpt --help
ggpt is a CLI tool to interact with OpenAI's GPT language model. ggpt wraps OpenAI's completion feature, via their API, outputting the result directly in the terminal.

Usage:
  ggpt [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  configure   Configure your OpenAI API key
  help        Help about any command
  last        Returns the output of the previous query.
  prompt      Call GPT autocomplete with the given string as a prompt

Flags:
  -h, --help     help for ggpt
  -t, --toggle   Help message for toggle

Use "ggpt [command] --help" for more information about a command.
```
## Usage
todo

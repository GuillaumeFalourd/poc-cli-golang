# poc-cli-golang

POC of a CLI project using Golang 🐿

## Requirements

- Golang installed (go1.15 or newer)
- Cobra library (`go get -u github.com/spf13/cobra/cobra`)
- Any code editor of your choice

## Setup & Run

Then at the repository `root` run: `go install my-cli`

Then, the `my-cli` command should be recognized on your terminal.

_Note: If not, run `export PATH=$PATH:$(go env GOPATH)/bin` or check your `$HOME/go/bin directory` if the `my-cli` bin is there_

```bash
➜  poc-cli-golang git:(main) ✗ my-cli
Hello CLI
```

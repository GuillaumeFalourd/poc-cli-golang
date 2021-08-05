# poc-cli-golang

POC of a CLI project using Golang inspired from [this article](https://towardsdatascience.com/how-to-create-a-cli-in-golang-with-cobra-d729641c7177).

## Requirements

- Golang installed (go1.15 or newer)
- Cobra library (`go get -u github.com/spf13/cobra/cobra`)
- Any code editor of your choice

## Setup & Update

At the repository `root` run: `go install my-cli`

Then, the `my-cli` command should be recognized on your terminal.

_Note: If not, run `export PATH=$PATH:$(go env GOPATH)/bin` or check your `$HOME/go/bin directory` if the `my-cli` bin is there_

## Run

**Available Commands:**

### Login

Will return SUCCESS if `username` and `password` flags are the same.

```bash
my-cli login -u="username" -p = "password"
```

### Add

Will return the sum of all numbers used as arguments.

```bash
my-cli add 1 2 3
```

#### Add Even (Add sub-command)

Will return the sum of all even numbers used as arguments.

```bash
my-cli add even 1 2 3
```

#### Add Odd (Add sub-command)

Will return the sum of all odd numbers used as arguments.

```bash
my-cli add odd 1 2 3
```

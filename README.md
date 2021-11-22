# poc-cli-golang

POC of a CLI project using Golang inspired from [this article](https://towardsdatascience.com/how-to-create-a-cli-in-golang-with-cobra-d729641c7177).

## 📑 Requirements

- Golang installed (go1.15 or newer)
- Cobra library (`go get -u github.com/spf13/cobra/cobra`)
- Any code editor of your choice

## ⚙️ Setup & Update

At the repository `root` run:

```bash
go install my-cli
```

Then, the `my-cli` command should be recognized on your terminal.

_**Note**: If `my-cli` is not recognized by the terminal, please run:_

```bash
export PATH=$PATH:$(go env GOPATH)/bin
```

_...or check your `$HOME/go/bin directory` if the `my-cli` bin is there)._

## 🚀 Run

### Login

Will return SUCCESS if `username` and `password` flags are the same.

```bash
my-cli login -u="username" -p = "password"
```

### Add

Will return the sum of all numbers used as arguments.

```bash
my-cli add 1 2 3 4
```

#### Add Even (Add sub-command)

Will return the sum of all even numbers used as arguments.

```bash
my-cli add even 1 2 3 4
```

#### Add Odd (Add sub-command)

Will return the sum of all odd numbers used as arguments.

```bash
my-cli add odd 1 2 3 4
```

### Create

Will create something with the specified **arg** `<name>` using the **flag** `--env-from`:

```bash
my-cli create <name> --env-from="dev"
```
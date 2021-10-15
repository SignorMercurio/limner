<p align="center">
<img src="img/logo.png" width="256" />
</p>

<h1 align="center">Limner</h1>

<div align="center">

Limner colorizes and transforms CLI outputs.

<a href="https://github.com/SignorMercurio/limner/actions"><img src="https://img.shields.io/github/workflow/status/SignorMercurio/limner/Go?logo=GitHub" /></a>
<a href="https://codecov.io/gh/SignorMercurio/limner"><img src="https://codecov.io/gh/SignorMercurio/limner/branch/main/graph/badge.svg?token=PKWZK3BR9R"/></a>
<a href="https://goreportcard.com/report/github.com/SignorMercurio/limner"><img src="https://goreportcard.com/badge/github.com/SignorMercurio/limner" /></a>
<a href="https://github.com/SignorMercurio/limner/blob/main/LICENSE"> <img src="https://img.shields.io/github/license/SignorMercurio/limner" /></a>

</div>

<a href="https://asciinema.org/a/ZtR2TaQPJWSUwTSIInSKZmrFu" target="_blank"><img src="https://asciinema.org/a/ZtR2TaQPJWSUwTSIInSKZmrFu.svg" /></a>

## Motivation

When playing with `kubectl`, I sometimes found it hard to extract the information I needed most from the extermely **long, mono-colored** output produced by the CLI. The same things happens when I'm using `curl` to test some REST APIs which usually return a long string of JSON.

Of course I can use GUI tools like *Kubernetes Dashboard* and *Postman*, but for simple operations that need to be performed swiftly, CLIs have their own advantages. Therefore, I made limner to bring some changes to the CLIs' output.

## Installation

### Download a release binary
Go to [Release Page](https://github.com/SignorMercurio/limner/releases), download a release and run:
```bash
tar zxvf lm_[version]_[os]_[arch].tar.gz
cd lm_[version]_[os]_[arch]
mv ./lm_[os]_[arch] ./lm
chmod +x ./lm
[your command] | ./lm
```

Remember to replace the text in `[]`.

### Manual Installation

1. You'll need Go [installed](https://golang.org/doc/install).

2. Clone the repo:
```bash
git clone https://github.com/SignorMercurio/limner.git
cd limner
go build -o lm .
```

3. Run the command:
```bash
[your command] | ./lm
```

> Note: It's strongly recommended to add the binary to your $PATH, e.g. `/usr/local/bin`.

## Usage

### Basic Usage

In most cases, you don't need to append any arguments when using limner as it automatically detects the format of the output.

Colorize tables:

```bash
kubectl get po | lm
```

Colorize YAML files:

```bash
cat nginx-deploy.yml | lm
```

Colorize JSON responses:

```bash
curl https://api.github.com/users/SignorMercurio | lm
```

> TODO: Add support for more formats and transformation between different formats.

### Create a shortcut

Take `kubectl` as an example. 

#### Bash

Suppose you've already [configured autocompletion](https://kubernetes.io/docs/tasks/tools/included/optional-kubectl-configs-bash-linux/) for `kubectl` (Optional).

In your `.bash_profile` or `.bashrc`, append the following lines:
```bash
function k() {kubectl $@ | lm}
complete -o default -F __start_kubectl k
```

#### Zsh

Suppose you've already [configured autocompletion](https://kubernetes.io/docs/tasks/tools/included/optional-kubectl-configs-zsh/) for `kubectl` (Optional).

In your `.zprofile` or `.zshrc`, append the following lines:
```bash
function k() {kubectl $@ | lm}
compdef k=kubectl
```

After the above steps, you'll be able to use `kubectl` with color and autocompletion like:
```bash
k get po
```

### Non-terminal output

When you choose to output the result to a file, or pass the result to other programs, through a pipe `|` or redirection `>`, you certainly do not want limner to colorize the output. The `--plain` flag (or `-p`) is meant for this, which prevent limner from colorizing the output.

### Customize color themes

You can use a config file to customize color themes. By default, limner will try to read `$HOME/.limner.yaml` but you can specify the config file with `-c`, for example:

```bash
kubectl get po | lm -c config/limner.yml
```

And here's an example of config file, which uses the same color theme as the default one:

```yaml
key_color: Red
string_color: Green
bool_color: Yellow
number_color: Yellow
null_color: Cyan
header_color: Blue
column_colors:
  - White
  - Cyan
```

Possible colors include `Red`, `Green`, `Yellow`, `Cyan`, `Blue`, `Magenta`, `White` and `Black`.

### Enforce types on the output

Specify `-t` to force limner to view the output as a specific type: YAML / JSON / table, etc. For example:

```bash
kubectl describe deploy/nginx | lm -t yaml
```

> Note: Specifying `-t yaml` in `kubectl describe` is not necessary.

## Contributions

Any contributions are welcome. Please feel free to:

- Open an Issue
- Creating a Pull Request
- Comment in an Issue / PR
- Open a Discussion

Thank you for willing to contribute to this project!

## Roadmap

- [x] Basic colorization
  - [x] YAML
  - [x] JSON
  - [x] Tables
  - [ ] ...
- [ ] Simple data format transformation
  - [ ] YAML <-> JSON
  - [ ] ...

If you have any suggestions for the project, please don't hesitate to open an [issue](https://github.com/SignorMercurio/limner/issues) or [pull request](https://github.com/SignorMercurio/limner/pulls).

## LICENSE

Distributed under the MIT License. See [LICENSE](LICENSE) for more information.

## Acknowledgements

Inspired by the following incredible projects:

- [cobra](https://github.com/spf13/cobra)
- [kubecolor](https://github.com/dty1er/kubecolor)
- [yh](https://github.com/andreazorzetto/yh)

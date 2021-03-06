<p align="center">
<img src="img/logo.png" width="256" />
</p>

<h1 align="center">Limner</h1>

<div align="center">

Limner colorizes and transforms CLI outputs.

English | [简体中文](README_zh.md)

<a href="https://github.com/SignorMercurio/limner/actions"><img src="https://img.shields.io/github/workflow/status/SignorMercurio/limner/Go?logo=GitHub" /></a>
<a href="https://codecov.io/gh/SignorMercurio/limner"><img src="https://codecov.io/gh/SignorMercurio/limner/branch/main/graph/badge.svg?token=PKWZK3BR9R"/></a>
<a href="https://goreportcard.com/report/github.com/SignorMercurio/limner"><img src="https://goreportcard.com/badge/github.com/SignorMercurio/limner" /></a>
<a href="https://github.com/SignorMercurio/limner/blob/main/LICENSE"> <img src="https://img.shields.io/github/license/SignorMercurio/limner" /></a>

<a href="https://asciinema.org/a/ZtR2TaQPJWSUwTSIInSKZmrFu" target="_blank"><img src="https://asciinema.org/a/ZtR2TaQPJWSUwTSIInSKZmrFu.svg" /></a>

</div>

## Motivation

When playing with `kubectl`, I sometimes found it hard to extract the information I needed most from the extermely **long, mono-colored** output produced by the CLI. The same things happens when I'm using `curl` to test some REST APIs which usually return a long string of JSON.

Of course I can use GUI tools like _Kubernetes Dashboard_ and _Postman_, but for simple operations that need to be performed swiftly, CLIs have their own advantages. Therefore, I made limner to bring some changes to the CLIs' output.

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

#### Colorize tables

```bash
kubectl get po | lm
```

![Table colorization](img/table.png)

#### Colorize YAML files

```bash
kubectl describe deploy/nginx | lm
```

![YAML colorization](img/yml.png)

#### Colorize JSON responses

```bash
curl -s https://api.github.com/users/SignorMercurio | lm
```

![JSON colorization](img/json.png)

#### Transform YAML to JSON

```bash
cat nginx/deploy.yml | lm tr -i yaml -o json
```

You can always omit `-i yaml` as long as the format of input is YAML (or looks like YAML). The same is true for other formats.

![YAML->JSON transformation](img/yml2json.png)

#### Transform JSON to YAML

```bash
curl -s https://jsonplaceholder.typicode.com/users/1/albums | lm tr -o yml
```

![JSON->YAML transformation](img/json2yml.png)

> Note: Limner is only designed for dealing with outputs. Do not use it with commands that need to receive input from stdin.

> TODO: Add support for more formats and transformation between different formats.

### Create a shortcut

Take `kubectl` as an example.

#### Bash

Suppose you've already [configured autocompletion](https://kubernetes.io/docs/tasks/tools/included/optional-kubectl-configs-bash-linux/) for `kubectl` (Optional).

In your `.bash_profile` or `.bashrc`, append the following lines:

```bash
function k() {kubectl $@ | lm}
complete -o default -F __start_kubectl k
alias kx="kubectl exec -it"
```

#### Zsh

Suppose you've already [configured autocompletion](https://kubernetes.io/docs/tasks/tools/included/optional-kubectl-configs-zsh/) for `kubectl` (Optional).

In your `.zprofile` or `.zshrc`, append the following lines:

```bash
function k() {kubectl $@ | lm}
compdef k=kubectl
alias kx="kubectl exec -it"
```

After the above steps, you'll be able to use `kubectl` with color and autocompletion like:

![Using a shortcut](img/shortcut.png)

### Non-terminal output

When you choose to output the result to a file, or pass the result to other programs, through a pipe `|` or redirection `>`, you certainly do not want limner to colorize the output. The `--plain` flag (or `-p`) is meant for this, which prevent limner from colorizing the output.

### Customize color themes

You can use a config file to customize color themes. By default, limner will try to read `$HOME/.limner.yaml` but you can specify the config file with `-c`, for example:

```bash
kubectl get po | lm -c config/limner.yml
```

The default config file looks like:

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

And here's an example of config file, which is suitable for a light-background terminal:

```yaml
key_color: Blue
string_color: Green
bool_color: Red
number_color: Red
null_color: Cyan
header_color: Magenta
column_colors:
  - Black
  - Cyan
```

Possible colors include `Red`, `Green`, `Yellow`, `Cyan`, `Blue`, `Magenta`, `White` and `Black`.

### Enforce types on the output

Specify `-t` to force limner to view the output as a specific type: YAML / JSON / table, etc. For example:

```bash
kubectl describe deploy/nginx | lm -t yaml
```

> Note: Specifying `-t yaml` in `kubectl describe` is not necessary.

### Tranformation

As you can see from the section [Transform YAML to JSON](#transform-yaml-to-json), all you need is:

```bash
[something of input type] | lm -o [output type]
```

Use `-i [input type]` if you want to force the limner to view the input as a specific type.

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
- [x] Simple data format transformation
  - [x] YAML <-> JSON
  - [ ] ...

If you have any suggestions for the project, please don't hesitate to open an [issue](https://github.com/SignorMercurio/limner/issues) or [pull request](https://github.com/SignorMercurio/limner/pulls).

## LICENSE

Distributed under the MIT License. See [LICENSE](LICENSE) for more information.

## Acknowledgements

Inspired by the following incredible projects:

- [cobra](https://github.com/spf13/cobra)
- [kubecolor](https://github.com/dty1er/kubecolor)
- [yh](https://github.com/andreazorzetto/yh)

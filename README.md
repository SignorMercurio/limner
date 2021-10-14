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

## Installation

// TODO

## Usage

### Basic Usage

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
curl https://api.github.com/SignorMercurio | lm
```

// TODO

### Non-terminal output

When you choose to output the result to a file, or pass the result to other programs, through a pipe `|` or redirection `>`, you certainly do not want limner to colorize the output. The `--plain` flag (or `-p`) is meant for this, which prevent limner from colorizing the output.

### Customize color themes

You can use a config file to customize color themes. By default, limner will try to read `$HOME/.limner.yaml` but you can specify the config file with `-c`, for example:

```bash
k get po | lm -c config/limner.yml
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
k describe deploy/nginx | lm -t yaml
```

> You don't actually need to use the flag at most of the time because limner automatically detects the possible format of the output.

// TODO

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

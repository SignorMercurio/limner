# limner

Limner colorizes and transforms CLI outputs.

[![asciicast](https://asciinema.org/a/ZtR2TaQPJWSUwTSIInSKZmrFu.svg)](https://asciinema.org/a/ZtR2TaQPJWSUwTSIInSKZmrFu)

## Installation

// TODO

## Usage

### Basic Usage

Colorize tables:

```bash
kubectl get po|lm
```

Colorize YAML files:

```bash
cat nginx-deploy.yml|lm
```

Colorize JSON responses:

```bash
curl https://api.github.com/SignorMercurio|lm
```

// TODO

### Non-terminal output

When you choose to output the result to a file, or pass the result to other programs, through a pipe `|` or redirection `>`, you certainly do not want limner to colorize the output. The `--plain` flag is meant for this, which prevent limner from colorizing the output.

### Custom color themes

// TODO

### Enforce types on the output

Specify `-t` to force limner to view the output as a specific type: YAML / JSON / XML / table, etc. For example:

```bash
k describe deploy/nginx | lm -t yml
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

## LICENSE

Distributed under the MIT License. See [LICENSE](LICENSE) for more information.

## Acknowledgements

Inspired by the following incredible projects:

- [cobra](https://github.com/spf13/cobra)
- [kubecolor](https://github.com/dty1er/kubecolor)
- [yh](https://github.com/andreazorzetto/yh)

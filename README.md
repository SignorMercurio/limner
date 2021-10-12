# limner

Limner colorizes and transforms CLI outputs.

// TODO

## Installation

// TODO

## Usage

### Basic Usage

```bash
lm -- kubectl get po
```

```bash
lm -- cat nginx-deploy.yml
```

```bash
lm -- curl https://api.github.com/
```

// TODO

### Create an alias (Recommended)

Take `kubectl` as an example. In your `.bash_profile`, `.zprofile`, etc., append the following:

```bash
alias k="lm -- kubectl"
```

Then you'll be able to use `kubectl` like normal, with the ability provided by limner:

```bash
k get po
```

// TODO

### Non-terminal output

When you choose to output the result to a file, or pass the result to other programs, through a pipe `|` or redirection `>`, you certainly do not want limner to colorize the output. The `--plain` flag is meant for this, which prevent limner from colorizing the output.

### On light backgrounds

Specify `--light-bg` to adapt a more suitable color theme in a terminal with light background

### Enforce types on the output

Specify `-t` to force limner to view the output as a specific type: YAML / JSON / XML / table, etc. For example:

```bash
lm -t yml k describe deploy/nginx
```

// TODO

## Contributions

Any contributions are welcome. Please feel free to:

- Open an Issue
- Creating a Pull Request
- Comment in an Issue / PR
- Open a Discussion

Thank you for willing to contribute to this project!

## LICENSE

MIT

## Acknowledgements

Inspired by the following incredible projects:

- [kubecolor](https://github.com/dty1er/kubecolor)
- [yh](https://github.com/andreazorzetto/yh)

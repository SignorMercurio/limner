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

### Autocompletion

To utilize the autocompetion for the original command, you may need to apply some changes to your shell config files.

#### Bash

Take `kubectl` as an example. In your `.bashrc`, append the following:

```bash
complete -o default -F __start_kubectl limner
```

#### Zsh

Take `kubectl` as an example. In your `.zshrc`, append the following:

```bash
source <(kubectl completion zsh)
compdef lm=kubectl
```

If you encounter problems like `command not found: compdef`, you may need to write it like:

```bash
autoload -Uz compinit; compinit
source <(kubectl completion zsh)
compdef lm=kubectl
```

### Non-terminal output

When you choose to output the result to a file, or pass the result to other programs, through a pipe `|` or redirection `>`, you certainly do not want limner to colorize the output. The `--plain` flag is meant for this, which prevent limner from colorizing the output.

### Custom color themes

// TODO

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

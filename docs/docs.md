# gosync

## Usage
> This tool will copy files really fast

gosync

## Flags
|Flag|Usage|
|----|-----|
|`--debug`|enable debug messages|
|`--disable-update-checks`|disables update checks|
|`--raw`|print unstyled raw output (set it if output is written to a file)|

## Commands
|Command|Usage|
|-------|-----|
|`gosync completion`|Generate the autocompletion script for the specified shell|
|`gosync help`|Help about any command|
# ... completion
`gosync completion`

## Usage
> Generate the autocompletion script for the specified shell

gosync completion

## Description

```
Generate the autocompletion script for gosync for the specified shell.
See each sub-command's help for details on how to use the generated script.

```

## Commands
|Command|Usage|
|-------|-----|
|`gosync completion bash`|Generate the autocompletion script for bash|
|`gosync completion fish`|Generate the autocompletion script for fish|
|`gosync completion powershell`|Generate the autocompletion script for powershell|
|`gosync completion zsh`|Generate the autocompletion script for zsh|
# ... completion bash
`gosync completion bash`

## Usage
> Generate the autocompletion script for bash

gosync completion bash

## Description

```
Generate the autocompletion script for the bash shell.

This script depends on the 'bash-completion' package.
If it is not installed already, you can install it via your OS's package manager.

To load completions in your current shell session:

	source <(gosync completion bash)

To load completions for every new session, execute once:

#### Linux:

	gosync completion bash > /etc/bash_completion.d/gosync

#### macOS:

	gosync completion bash > $(brew --prefix)/etc/bash_completion.d/gosync

You will need to start a new shell for this setup to take effect.

```

## Flags
|Flag|Usage|
|----|-----|
|`--no-descriptions`|disable completion descriptions|
# ... completion fish
`gosync completion fish`

## Usage
> Generate the autocompletion script for fish

gosync completion fish

## Description

```
Generate the autocompletion script for the fish shell.

To load completions in your current shell session:

	gosync completion fish | source

To load completions for every new session, execute once:

	gosync completion fish > ~/.config/fish/completions/gosync.fish

You will need to start a new shell for this setup to take effect.

```

## Flags
|Flag|Usage|
|----|-----|
|`--no-descriptions`|disable completion descriptions|
# ... completion powershell
`gosync completion powershell`

## Usage
> Generate the autocompletion script for powershell

gosync completion powershell

## Description

```
Generate the autocompletion script for powershell.

To load completions in your current shell session:

	gosync completion powershell | Out-String | Invoke-Expression

To load completions for every new session, add the output of the above command
to your powershell profile.

```

## Flags
|Flag|Usage|
|----|-----|
|`--no-descriptions`|disable completion descriptions|
# ... completion zsh
`gosync completion zsh`

## Usage
> Generate the autocompletion script for zsh

gosync completion zsh

## Description

```
Generate the autocompletion script for the zsh shell.

If shell completion is not already enabled in your environment you will need
to enable it.  You can execute the following once:

	echo "autoload -U compinit; compinit" >> ~/.zshrc

To load completions in your current shell session:

	source <(gosync completion zsh); compdef _gosync gosync

To load completions for every new session, execute once:

#### Linux:

	gosync completion zsh > "${fpath[1]}/_gosync"

#### macOS:

	gosync completion zsh > $(brew --prefix)/share/zsh/site-functions/_gosync

You will need to start a new shell for this setup to take effect.

```

## Flags
|Flag|Usage|
|----|-----|
|`--no-descriptions`|disable completion descriptions|
# ... help
`gosync help`

## Usage
> Help about any command

gosync help [command]

## Description

```
Help provides help for any command in the application.
Simply type gosync help [path to command] for full details.
```


---
> **Documentation automatically generated with [PTerm](https://github.com/pterm/cli-template) on 29 October 2022**

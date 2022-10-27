# gosync

## Usage
> This cli template shows the date and time in the terminal

gosync

## Description

```
This is a template CLI application, which can be used as a boilerplate for awesome CLI tools written in Go.
This template prints the date or time to the terminal.
```
## Examples

```bash
cli-template date
cli-template date --format 20060102
cli-template time
cli-template time --live
```

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
|`gosync date`|Prints the current date.|
|`gosync help`|Help about any command|
|`gosync time`|Prints the current time|
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
# ... date
`gosync date`

## Usage
> Prints the current date.

gosync date

## Flags
|Flag|Usage|
|----|-----|
|`-f, --format string`|specify a custom date format (default "02 Jan 06")|
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
# ... time
`gosync time`

## Usage
> Prints the current time

gosync time

## Description

```
You can print a live clock with the '--live' flag!
```

## Flags
|Flag|Usage|
|----|-----|
|`-l, --live`|live output|


---
> **Documentation automatically generated with [PTerm](https://github.com/pterm/cli-template) on 27 October 2022**

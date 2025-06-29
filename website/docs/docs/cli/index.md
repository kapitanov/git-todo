# Command Line Interface

## Overview

`git-todo` provides a 100% functional command line interface (CLI) for managing your TODO items.
Anything you can do in the TUI - you can also do via the command line.
Besides, certain operations are only available through the CLI, such as installing git hooks.

All commands start with `git todo`, followed by the specific action you want to perform.
Alternatively, you can use the full name `git-todo` if you prefer.

## Text-mode user interface (TUI)

The TUI interface provides a user-friendly way to manage your TODO items.
You can navigate through the list of TODOs, check or uncheck items, edit their details, and remove them as needed.

To open the TUI, simply run:

```bash
git todo
```

This will display a list of your TODO items, allowing you to interact with them using keyboard shortcuts.

[Read more about the TUI interface :material-arrow-right: ](../tui.md){ .md-button }

## Command Reference

Here's a quick reference for the available commands:

| Command                          | Description                                |
| -------------------------------- | ------------------------------------------ |
| [`git todo`](../tui.md)          | Open the TUI interface for managing TODOs. |
| [`git todo add`](add.md)         | Add a new TODO item.                       |
| [`git todo check`](check.md)     | Mark a TODO item as "completed".           |
| [`git todo clear`](clear.md)     | Remove all TODO items.                     |
| [`git todo deinit`](deinit.md)   | Uninstall git hooks for TODO management.   |
| [`git todo edit`](edit.md)       | Edit a TODO item.                          |
| [`git todo init`](init.md)       | Install git hooks for TODO management.     |
| [`git todo ls`](ls.md)           | List all TODO items.                       |
| [`git todo rm`](rm.md)           | Remove a TODO item.                        |
| [`git todo uncheck`](uncheck.md) | Mark a TODO item as "incomplete".          |
| [`git todo view`](view.md)       | View a TODO item.                          |

## Using `git-todo` in scripts

You can use `git-todo` commands in your scripts to automate TODO management.
For example, you can create a script that adds a TODO item automatically:

```bash
#!/bin/bash

git todo add "Automate TODO management"
```

By default, all `git-todo` commands will produce a human-readable output.
But all commands also support the `--quiet` flag, which suppresses the output - but you still will have exit codes and stdout print.
Please refer to a specific command's documentation for more details on its usage and options.

Note that the `--quiet` flag disabled all interactive prompts - both [`EDITOR`](#editor-support)-based and prompt-based.
Also, running the [`git todo ls -q`](./ls.md) command will produce a different, script-friendly output instead of a pretty one.

We encourage you to use `git-todo` in your scripts to streamline your workflow and keep track of your tasks efficiently!
And if you would like to share your ways of using `git-todo` in
scripts - [you are more than welcome!](https://github.com/kapitanov/git-todo/issues/new?label=feedback)

### Exit codes

Regardless of the `--quiet` flag, all `git-todo` commands use this exit codes:

## Exit codes

| Exit Code | Description                                                                                  |
| --------: | -------------------------------------------------------------------------------------------- |
|       `0` | The TODO item was successfully added.                                                        |
|     `128` | Current directory is not a git repository.                                                   |
|       `1` | An attempt was made to edit or remove a non-existing item                                    |
|       `2` | A TODO item with the same title already exists (and no `-u`/`--unless-exists` flag was set). |
|       `3` | An operation has been canceled (via `SIGINT` signal).                                        |
|       `9` | An error occurred while adding the TODO item.                                                |

## `EDITOR` support

Certain commands like `git todo add` and `git todo edit` can open an editor for you to input or modify TODO items.
You can set your preferred editor by configuring the `EDITOR` environment variable. For example, to use `nano`, you can run:

```bash
export EDITOR=nano
```

This will ensure that when you run commands that require an editor, it will open in `nano`.
You can replace `nano` with any text editor of your choice, such as `vim`, `code`, or `emacs`.

By default, if the `EDITOR` variable is not set, `git-todo` will try to use a default editor:

- `nano` on macOS;
- `vi` on Linux;
- `notepad` on Windows.

## Item IDs and selectors

Each TODO item has a unique ID that can be used to reference it in commands.
You can use the ID to perform operations like editing or removing a specific TODO item, e.g.:

```bash
git todo view e885a108
```

Besides, you can use a short version of the ID - just type the first few characters of the ID.
If the short ID is unique, it will be resolved to the full ID automatically.
But if there are multiple items with the same short ID,
you will get an "item not found" error.

```bash
git todo view e88
```

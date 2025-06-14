# `git todo init`

This command installs `git-todo`-powered git hooks for the current git repository.

It's an antipode of the [`git todo deinit`](./deinit.md) command.

## How it works

When you run `git todo init`, it will locate the root of the current git repository
and install the necessary git hooks into the `.git/hooks` directory.
The hooks will be set up to run `git-todo` commands at appropriate times, such as before committing or pushing changes.

IF there are any existing hooks in the `.git/hooks` directory, they will be overwritten by the new hooks installed by `git-todo`,
but in a smart way:

- If an existing git hook already contains a `git-todo` command, it will be left as is.
- Otherwise:
    - The existing hook will be backed up to a file with the `.bak` extension.
    - The necessary `git-todo` command will be appended into the git hook.

!!! note

    For more information about `git-todo` hooks, please refer to [the corresponding section](../git-hooks.md) of the documentation.

## Usage

```bash
git todo init [-q|--quiet] [-v|--verbose]
```

## Flags and arguments

| Flag/Argument     | Mandatory | Description                                                                                      |
| ----------------- | --------- | ------------------------------------------------------------------------------------------------ |
| `-q`, `--quiet`   | No        | Suppress output, only exit codes will be printed.                                                |
| `-v`, `--verbose` | No        | Print additional information about the operation.                                                |
| `-f`, `--force`   | No        | Force githooks installation, even if hooks are already installed (will overwrite existing hooks) |

## Examples

```bash
# Install git-todo hooks for the current git repository.
$ git todo init
Git hooks installed successfully.

# Re-install git-todo hooks for the current git repository.
$ git todo init -f
Git hooks installed successfully.

# Install git-todo hooks for the current git repository with verbose output.
$ git todo init -v
11:58PM DBG executing command cmd="/opt/homebrew/bin/git git rev-parse --show-toplevel"
11:58PM INF discovered git repository root root=/Users/username/git-repository
11:58PM DBG loaded model file path=/Users/username/git-repository/.git/TODO
11:58PM INF installing git hook file=/Users/username/git-repository/.git/hooks/pre-commit hook=pre-commit
11:58PM INF git hook is already installed hook=pre-commit
11:58PM DBG set git hook as executable file=/Users/username/git-repository/.git/hooks/pre-commit hook=pre-commit perm=493
11:58PM INF installing git hook file=/Users/username/git-repository/.git/hooks/pre-push hook=pre-push
11:58PM INF git hook is already installed hook=pre-push
11:58PM DBG set git hook as executable file=/Users/username/git-repository/.git/hooks/pre-push hook=pre-push perm=493
Git hooks installed successfully
```

### Scripting usage examples

For scripting, you should use `-q` or `--quiet` flag to suppress output and avoid cluttering the console.

```bash
# Install git-todo hooks
$ git todo init -qf
```

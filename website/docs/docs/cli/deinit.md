# `git todo deinit`

This command uninstalls `git-todo`-powered git hooks from the current git repository.

It's an antipode of the [`git todo init`](./init.md) command.

## How it works

When you run `git todo deinit`, it will locate the root of the current git repository
and install the necessary git hooks into the `.git/hooks` directory.
Then, all `git-todo` hooks will be removed from the repository.

IF there are any existing hooks in the `.git/hooks` directory, they will be overwritten by the new hooks installed by `git-todo`,
but in a smart way:

- If an existing git hook doesn't contain a `git-todo` command, it will be left as is.
- Otherwise:
  - The existing hook will be backed up to a file with the `.bak` extension.
  - The `git-todo` command will be removed fro the git hook.

## Usage

```bash
git todo deinit [-q|--quiet] [-v|--verbose]
```

## Flags and arguments

| Flag/Argument     | Mandatory | Description                                                                                    |
| ----------------- | --------- | ---------------------------------------------------------------------------------------------- |
| `-q`, `--quiet`   | No        | Suppress output, only exit codes will be printed.                                              |
| `-v`, `--verbose` | No        | Print additional information about the operation.                                              |
| `-f`, `--force`   | No        | Force githooks uninstallation, even if hooks are not installed (will overwrite existing hooks) |

## Examples

```bash
# Uninstall git-todo hooks from the current git repository.
$ git todo deinit
Git hooks uninstalled successfully.

# Forcibly uninstall git-todo hooks from the current git repository.
$ git todo deinit -f
Git hooks uninstalled successfully.

# Uninstall git-todo hooks from the current git repository with verbose output.
$ git todo deinit -v
12:10AM DBG executing command cmd="/opt/homebrew/bin/git git rev-parse --show-toplevel"
12:10AM INF discovered git repository root root=/Users/username/git-repository
12:10AM DBG loaded model file path=/Users/username/git-repository/.git/TODO
12:10AM INF uninstalling git hook file=/Users/username/git-repository/.git/hooks/pre-commit hook=pre-commit
12:10AM INF created a backup of a git hook file=/Users/username/git-repository/.git/hooks/pre-commit.bak hook=pre-commit
12:10AM INF updated git hook hook=pre-commit
12:10AM DBG set git hook as executable file=/Users/username/git-repository/.git/hooks/pre-commit hook=pre-commit perm=493
12:10AM INF uninstalling git hook file=/Users/username/git-repository/.git/hooks/pre-push hook=pre-push
12:10AM INF created a backup of a git hook file=/Users/username/git-repository/.git/hooks/pre-push.bak hook=pre-push
12:10AM INF updated git hook hook=pre-push
12:10AM DBG set git hook as executable file=/Users/username/git-repository/.git/hooks/pre-push hook=pre-push perm=493
Git hooks uninstalled successfully.
```

### Scripting usage examples

For scripting, you should use `-q` or `--quiet` flag to suppress output and avoid cluttering the console.

```bash
# Uninstall git-todo hooks
$ git todo deinit -qf
```

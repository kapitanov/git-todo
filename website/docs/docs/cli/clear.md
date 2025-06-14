# `git todo clear`

This command allows you to permanently remove all TODO items.

## Usage

```bash
git todo clear [-q|--quiet] [-v|--verbose]
```

## Flags and arguments

| Flag/Argument     | Mandatory | Description                                         |
| ----------------- | --------- | --------------------------------------------------- |
| `-q`, `--quiet`   | No        | Suppress output, only exit codes will be printed.   |
| `-v`, `--verbose` | No        | Print additional information about the operation.   |
| `-f`, `--force`   | No        | Disables interactive confirmation.                  |

## Examples

```bash
# Remove all TODO items, but the removal is not confirmed.
$ git todo clear
Are you sure you want to remove all TODO items (y/n)? n

# Remove all TODO items, and the removal is confirmed.
$ git todo clear
Are you sure you want to remove all TODO items (y/n)? y
All TODO items have been deleted

# Remove all TODO items without asking for confirmation - with more verbose output.
$ git todo clear -vf
1:54PM DBG executing command cmd="/opt/homebrew/bin/git git rev-parse --show-toplevel"
11:54PM INF discovered git repository root root=/Users/username/git-repository
11:54PM DBG loaded model file path=/Users/username/git-repository/.git/TODO
All TODO items have been deleted
```

### Scripting usage examples

For scripting, you should use `-q` or `--quiet` flag to suppress output and avoid cluttering the console.
Also, you should use `-f` or `--force` to skip confirmation prompts, as they won't be available for scripted scenarios.

```bash
# Remove all TODO items
$ git todo clear -qf
```

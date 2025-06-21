# `git todo rm`

This command allows you to permanently remove a TODO item (or few items).

## Usage

```bash
git todo check [-q|--quiet] [-v|--verbose] [<id>...]
```

## Flags and arguments

| Flag/Argument     | Mandatory | Description                                       |
| ----------------- | --------- | ------------------------------------------------- |
| `-q`, `--quiet`   | No        | Suppress output, only exit codes will be printed. |
| `-v`, `--verbose` | No        | Print additional information about the operation. |
| `-f`, `--force`   | No        | Disables interactive confirmation.                |
| `<id>...`         | Yes       | An ID of a TODO item to be removed.               |

## Examples

```bash
# Remove a TODO item [e885a108]
$ git todo e885a108
Are you sure you want to remove TODO item [e885a108] "Write some code" (y/n)? y
TODO item [e885a108] "Write some code" has been removed

# Remove TODO items [e885a108] and [419ee57f], but the first removal is not confirmed.
$ git todo rm e885a108 419ee57f
Are you sure you want to remove TODO item [e885a108] "Write some code" (y/n)? n
Canceled removal of TODO item [e885a108] "Write some code"
Are you sure you want to remove TODO item [419ee57f] "Write some useful documentation as it is important" (y/n)? y
TODO item [419ee57f] "Write some useful documentation as it is important" has been removed

# Remove TODO items [e885a108] and [419ee57f] without asking for confirmation - with more verbose output.
$ git todo rm -vf e885a108 419ee57f
11:50PM DBG executing command cmd="/opt/homebrew/bin/git git rev-parse --show-toplevel"
11:50PM INF discovered git repository root root=/Users/username/git-repository
11:50PM DBG loaded model file path=/Users/username/git-repository/.git/TODO
TODO item [e885a108] "Write some code" has been removed
TODO item [419ee57f] "Write some useful documentation as it is important" has been removed
```

### Scripting usage examples

For scripting, you should use `-q` or `--quiet` flag to suppress output and avoid cluttering the console.
Also, you should use `-f` or `--force` to skip confirmation prompts, as they won't be available for scripted scenarios.

```bash
# Remove a TODO item [e885a108]
$ git todo rm -qf e885a108
e885a108

# Remove TODO items [e885a108] and [419ee57f]
$ git todo rm -qf e885a108 419ee57f
e885a108
419ee57f
```

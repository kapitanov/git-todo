# `git todo check`

This command allows you to mark a TODO item (or few items) as **completed**.

It's an antipode of the [`git todo uncheck`](./uncheck.md) command.

## Usage

```bash
git todo check [-q|--quiet] [-v|--verbose] [<id>...]
```

## Flags and arguments

| Flag/Argument     | Mandatory | Description                                         |
| ----------------- | --------- | --------------------------------------------------- |
| `-q`, `--quiet`   | No        | Suppress output, only exit codes will be printed.   |
| `-v`, `--verbose` | No        | Print additional information about the operation.   |
| `<id>...`         | Yes       | An ID of a TODO item to be marked as **completed**. |

If an item is already marked as **completed**, no action will be taken.

## Examples

```bash
# Mark a TODO item [e885a108] as "completed"
$ git todo check e885a108
TODO item [e885a108] "Write some code" has been marked as completed

# Mark TODO items [e885a108] and [419ee57f] as "completed"
$ git todo check e885a108 419
TODO item [e885a108] "Write some code" is already marked as completed
TODO item [419ee57f] "Write some useful documentation as it is important" has been marked as completed

# Mark a TODO item [ae19ad18] as "completed" - with more verbose output.
$ git todo check -v ae
10:35PM DBG executing command cmd="/opt/homebrew/bin/git git rev-parse --show-toplevel"
10:35PM INF discovered git repository root root=/Users/username/git-repository
10:35PM DBG loaded model file path=/Users/username/git-repository/.git/TODO
TODO item [ae19ad18] "Build something great" has been marked as completed
```

### Scripting usage examples

For scripting, you should use `-q` or `--quiet` flag to suppress output and avoid cluttering the console.

```bash
# Mark a TODO item [e885a108] as "completed"
$ git todo check -q e885a108
e885a108

# Mark TODO items [e885a108] and [419ee57f] as "completed"
$ git todo check -q e885a108 419ee57f
419ee57f
```

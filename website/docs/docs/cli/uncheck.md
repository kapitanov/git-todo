# `git todo uncheck`

This command allows you to mark a TODO item (or few items) as **incomplete**.

It's an antipode of the [`git todo check`](./check.md) command.

## Usage

```bash
git todo uncheck [-q|--quiet] [-v|--verbose] [<id>...]
```

## Flags and arguments

| Flag/Argument     | Mandatory | Description                                          |
| ----------------- | --------- | ---------------------------------------------------- |
| `-q`, `--quiet`   | No        | Suppress output, only exit codes will be printed.    |
| `-v`, `--verbose` | No        | Print additional information about the operation.    |
| `<id>...`         | Yes       | An ID of a TODO item to be marked as **incomplete**. |

If an item is already marked as **incomplete**, no action will be taken.

## Examples

```bash
# Mark a TODO item [e885a108] as "incomplete"
$ git todo uncheck e885a108
TODO item [e885a108] "Write some code" has been marked as incomplete

# Mark TODO items [e885a108] and [419ee57f] as "incomplete"
$ git todo uncheck e88 419
TODO item [e885a108] "Write some code" is not marked as completed
TODO item [419ee57f] "Write some useful documentation as it is important" has been marked as incomplete

# Mark a TODO item [419ee57f] as "incomplete" - with more verbose output.
$ git todo uncheck -v 419ee57f
10:40PM DBG executing command cmd="/opt/homebrew/bin/git git rev-parse --show-toplevel"
10:40PM INF discovered git repository root root=/Users/username/git-repository
10:40PM DBG loaded model file path=/Users/username/git-repository/.git/TODO
TODO item[419ee57f] "Write some useful documentation as it is important" is not marked as completed
```

### Scripting usage examples

For scripting, you should use `-q` or `--quiet` flag to suppress output and avoid cluttering the console.

```bash
# Mark a TODO item [e885a108] as "incomplete"
$ git todo uncheck -q e885a108
e885a108

# Mark TODO items [e885a108] and [419ee57f] as "incomplete"
$ git todo uncheck -q e885a108 419ee57f
419ee57f
```

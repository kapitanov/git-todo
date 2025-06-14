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
# Mark a TODO item 1 as "incomplete"
$ git todo uncheck 1
TODO item #1 has been checked as incomplete (Write some code)

# Mark TODO items 1 and 3 as "incomplete"
$ git todo uncheck 1 3
TODO item 1 is not marked as completed (Write some code)
TODO item 3 has been marked as incomplete (Write some useful documentation as it is important)

# Mark a TODO item 4 as "incomplete" - with more verbose output.
$ git todo uncheck -v 4
10:40PM DBG executing command cmd="/opt/homebrew/bin/git git rev-parse --show-toplevel"
10:40PM INF discovered git repository root root=/Users/username/git-repository
10:40PM DBG loaded model file path=/Users/username/git-repository/.git/TODO
TODO item 4 is not marked as completed (Build something great!)
```

### Scripting usage examples

For scripting, you should use `-q` or `--quiet` flag to suppress output and avoid cluttering the console.

```bash
# Mark a TODO item 1 as "incomplete"
$ git todo uncheck -q 1
1

# Mark TODO items 1 and 3 as "incomplete"
$ git todo uncheck -q 1 3
3
```

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
# Mark a TODO item 1 as "completed"
$ git todo check 1
TODO item 1 has been checked as completed (Write some code)

# Mark TODO items 1 and 3 as "completed"
$ git todo check 1 3
TODO item 1 is already marked as completed (Write some code)
TODO item 3 has been marked as completed (Write some useful documentation as it is important)

# Mark a TODO item 4 as "completed" - with more verbose output.
$ git todo check -v 4
10:35PM DBG executing command cmd="/opt/homebrew/bin/git git rev-parse --show-toplevel"
10:35PM INF discovered git repository root root=/Users/username/git-repository
10:35PM DBG loaded model file path=/Users/username/git-repository/.git/TODO
TODO item 4 has been marked as completed (Build something great!)
```

### Scripting usage examples

For scripting, you should use `-q` or `--quiet` flag to suppress output and avoid cluttering the console.

```bash
# Mark a TODO item 1 as "completed"
$ git todo check -q 1
1

# Mark TODO items 1 and 3 as "completed"
$ git todo check -q 1 3
3
```

# `git todo rm`

This command allows you to permanently remove a TODO item (or few items).

!!! note

    The removal of items affects their IDs, as IDs are basically table row numbers.
    So, removing an item 1 from a 3-items TODO list will have you with items 1 and 2, not 2 and 3:

    ```bash
    $ git todo ls
    1 · Write some code
    2 ✓ Write some tests as well
    3 · Write some useful documentation as it is important

    $ git todo rm -f 1
    TODO item 1 has been removed (Write some code)

    $ git todo ls
    1 ✓ Write some tests as well
    2 · Write some useful documentation as it is important
    ```

    However, removing few items at once won't be affected by this effect:
    `git-todo` will make sure that you won't need to recalculate the IDs manually before typing the command.

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
# Remove a TODO item 1
$ git todo rm 1
Are you sure you want to remove TODO item #1 "Write some code" (y/n)? y
TODO item 1 has been removed (Write some code)

# Remove TODO items 1 and 2, but the first removal is not confirmed.
$ git todo rm 1 2
Are you sure you want to remove TODO item #1 "Write some tests as well" (y/n)? n
Canceled removal of TODO item 1 "Write some tests as well"
Are you sure you want to remove TODO item #2 "Write some useful documentation as it is important" (y/n)? y
TODO item 2 has been removed (Write some useful documentation as it is important)

# Remove TODO items 1 and 2 without asking for confirmation - with more verbose output.
$ git todo rm -vf 1 2
11:50PM DBG executing command cmd="/opt/homebrew/bin/git git rev-parse --show-toplevel"
11:50PM INF discovered git repository root root=/Users/username/git-repository
11:50PM DBG loaded model file path=/Users/username/git-repository/.git/TODO
TODO item 1 has been removed (Write some tests as well)
TODO item 2 has been removed (Write some useful documentation as it is important)
```

### Scripting usage examples

For scripting, you should use `-q` or `--quiet` flag to suppress output and avoid cluttering the console.
Also, you should use `-f` or `--force` to skip confirmation prompts, as they won't be available for scripted scenarios.

```bash
# Remove a TODO item 1
$ git todo rm -qf 1
1

# Remove TODO items 1 and 2
$ git todo rm -qf 1 2
1
2
```

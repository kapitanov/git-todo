# `git todo edit`

This command allows you to edit a TODO item's title.

## Usage

```bash
git todo edit [-q|--quiet] [-v|--verbose] [-t|--title <title>] <id>
```

## Flags and arguments

| Flag/Argument                   | Mandatory | Description                                       |
| ------------------------------- | --------- | ------------------------------------------------- |
| `-q`, `--quiet`                 | No        | Suppress output, only exit codes will be printed. |
| `-v`, `--verbose`               | No        | Print additional information about the operation. |
| `-t <title>`, `--title <title>` | No        | The new title of the TODO item.                   |
| `<id>`                          | Yes       | An ID of a TODO item to be edited.                |

If the `-t`/`--title` flag is provided, its value will be treated as the new title of the TODO item.
Otherwise, [an interactive editor](index.md#editor-support) will be opened to edit the title.
See [Examples](#examples) below for more details.

## Examples

```bash
# Rename an existing TODO item with the ID 1 to "Write some code"
$ git todo edit -t "Write some code" 1
TODO item 1 has been renamed:
  old: "Write some cde"
  new: "Write some code"

# Rename an existing TODO item with the ID 1 to "Write some code" again - no changes will be made
# However, the command will exit with a zero exit code
$ git todo edit -t "Write some code" 1
TODO item 1 has not been renamed: the new title is the same as the old one

# Open an interactive editor to type the new title of the TODO item.
# If you type any non-empty title and save the file, an ordinary rename will be carried out.
# If you don't provide -t/--title flag, this is the default behavior.
$ git todo edit 1
TODO item 1 has been renamed:
  old: "Write some code"
  new: "Write some code!"

# Rename an existing TODO item with the ID 1 to "Write some code" - with more verbose output
$ git todo edit -v -t "Write some code" 1
12:21AM DBG executing command cmd="/opt/homebrew/bin/git git rev-parse --show-toplevel"
12:21AM INF discovered git repository root root=/Users/username/git-repository
12:21AM DBG loaded model file path=/Users/username/git-repository/.git/TODO
TODO item 1 has been renamed:
  old: "Write some code!"
  new: "Write some code"
```


### Scripting usage examples

For scripting, you should use `-q` or `--quiet` flag to suppress output and avoid cluttering the console.
Also, you have to use `-t` or `--title` flag to specify the new title of the TODO item,
as the interactive prompt won't be available.

```bash
# Rename an existing TODO item with the ID 1 to "Write some code" - in a script
$ git todo edit -q -t "Write some code" 1
1

# Rename an existing TODO item with the ID 1 to "Write some code" again - no changes will be made
# However, the command will exit with a zero exit code
$ git todo edit -q -t "Write some code" 1
1
```

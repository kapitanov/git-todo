# `git todo add`

This command allows you to add a new TODO item to your list.
You can provide the title as arguments or edit it interactively.
A newely added TODO item will be marked as incomplete by default.

## Usage

```bash
git todo add [-q|--quiet] [-v|--verbose] [-u|--unless-exists] [<title>...]
```

## Flags and arguments

| Flag/Argument           | Mandatory | Description                                                                             |
| ----------------------- | --------- | --------------------------------------------------------------------------------------- |
| `-q`, `--quiet`         | No        | Suppress output, only exit codes will be printed.                                       |
| `-v`, `--verbose`       | No        | Print additional information about the operation.                                       |
| `-u`, `--unless-exists` | No        | Do not add the TODO item if it already exists in the list.                              |
| `<title>...`            | No        | The title of the TODO item to be added. If not provided, an editor will open for input. |

If any arguments are provided, they will be treated as the title of the new TODO item.
Otherwise, [an interactive editor](index.md#editor-support) will be opened to type the title.
See [Examples](#examples) below for more details.

Specify the `-u` or `--unless-exists` flag to make the tool "swallow" title conficts:
it will exit with `0` exit code even if an item with the same name already exists.
But even in this case a duplicate item will not be added to the list.
This flag is useful for scripts or automation where you want to avoid errors due to existing items.

## Examples

```bash
# Add a new incomplete TODO item with the title "Write some code"
$ git todo add "Write some code"
Added a new TODO item: 1 "Write some code"

# Add a new incomplete TODO item with the title "Write some tests as well"
# Note that the quotes are optional even if the title contains spaces!
$ git todo add Write some tests as well
Added a new TODO item: 2 "Write some tests as well"

# Add a new incomplete TODO item with the title "Write some useful documentation as it is important"
# Unfortunately, the quotes are omitted in this case
$ git todo add Write "some useful documentation" as it is important
Added a new TODO item: 3 "Write some useful documentation as it is important"

# Open an interactive editor to type the title of the new TODO item.
# If you type any non-empty title and save the file, it will be added to the list.
# If you don't provide any arguments, this is the default behavior.
$ git todo add
Added new TODO item: "Build something great!"

# Open an interactive editor to type the title of the new TODO item.
# It's same as above, but uses "nano" as the editor regardless of the default editor.
$ EDITOR=nano git todo add -v
10:29PM DBG executing command cmd="/opt/homebrew/bin/git git rev-parse --show-toplevel"
10:29PM INF discovered git repository root root=/Users/username/git-repository
10:29PM DBGloaded model file path=/Users/username/git-repository/.git/TODO
10:29PM DBG created temp file file=/var/folders/y6/8h5ky72d2gd9c_3m3qq_df8c0000gn/T/git-todo-2742934353.txt
10:29PM DBG running editor cmd="/usr/local/bin/code code --wait /var/folders/y6/8h5ky72d2gd9c_3m3qq_df8c0000gn/T/git-todo-2742934353.txt"
Added a new TODO item: 5 "And help people!"

# Trying to add an item that already exists.
$ git todo add "Write some code"
Added a new TODO item: 1 "Write some code"

# Trying to add an item that already exists - but now with the "--unless-exists".
$ git todo add -u "Write some code"
The TODO item already exists: 1 "Write some code"
```

### Scripting usage examples

For scripting, you should use `-q` or `--quiet` flag to suppress output and avoid cluttering the console.
Also, you can use `-u` or `--unless-exists` to avoid errors if the item already exists.

```bash
# Add a new TODO item with the title "Write some code" without output
$ git todo add -q "Write some code"
123

# Try adding a TODO item with the title "Write some code" while one already exists
$ git todo add -q "Write some code"
unable to create an item with title "Write some code": item already exists

# Try adding a TODO item with the title "Write some code" while one already exists
# Note that the `-u` flag is used to avoid errors if the item already exists - it would return an existing item ID instead
$ git todo add -qu "Write some code"
123
```

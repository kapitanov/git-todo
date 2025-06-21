# `git todo view`

This command allows you to print a TODO item (or few items).

It's a similar to the [`git todo ls`](./ls.md) command - it shares the same output format
but offers a different way to filter the TODO items table.

## Usage

```bash
git todo view [-q|--quiet] [-v|--verbose] [-j|--json] [-p|--plain] <id>...
```

## Flags and arguments

| Flag/Argument     | Mandatory | Description                                       |
| ----------------- | --------- | ------------------------------------------------- |
| `-q`, `--quiet`   | No        | Suppress output, only exit codes will be printed. |
| `-v`, `--verbose` | No        | Print additional information about the operation. |
| `-j`, `--json`    | No        | Print TODO items in JSON format.                  |
| `-p`, `--plain`   | No        | Print TODO items in the plain format.             |
| `<id>...`         | Yes       | An ID of a TODO item to be printed.               |


## Output formats

The command supports three output formats:

- **Default**: A human-readable colored format with a numbered list of TODO items:

    ```bash
    $ git todo view e885a108 4e3eeecc
    e885a108 ✓ Write some code
    4e3eeecc · Write some tests as well
    ```

- **JSON**: A structured JSON format, suitable for parsing in scripts:

    ```bash
    $ git todo view --json e885a108 4e3eeecc
    [
        {
            "id": "e885a108",
            "done": true,
            "title": "Write some code"
        },
        {
            "id": "4e3eeecc",
            "done": false,
            "title": "Write some tests as well"
        }
    ]
    ```

    The output schema is as follows:

    ```json
    [
        {
            "id"   : <string>,  // The ID of the TODO item
            "done" : <boolean>, // Whether the TODO item is completed (true) or incomplete (false)
            "title": <string>   // The title of the TODO item
        }
    ]
    ```

- **Plain**: A simple text format without any colors or additional formatting:

    ```bash
    $ git todo view -p e885a108 e885a108
    e885a108 DONE Write some code
    4e3eeecc TODO Write some tests as well
    ```

By default, the command will print the TODO items in a human-readable format.
But if the `-q` or `--quiet` flag is provided, it will force the plain-text format, unless its overriden by any other flags.

## Examples

```bash
# Print a single TODO item by its ID
$ git todo view 4e3eeecc
4e3eeecc · Write some tests as well

# Print a single TODO item by its ID in plain-text format
$ git todo view -p 4e3eeecc
4e3eeecc TODO Write some tests as well


# Print a single TODO item by its ID in JSON format
$ git todo view -j 4e3eeecc
[
    {
        "id": "4e3eeecc",
        "done": false,
        "title": "Write some tests as well"
    }
]

# Print multiple TODO items by their IDs in JSON format
$ git todo view -j 4e3eeecc 9612977c
[
    {
        "id": "4e3eeecc",
        "done": false,
        "title": "Write some tests as well"
    },
    {
        "id": "9612977c",
        "done": false,
        "title": "And help people!"
    }
]

# Print multiple TODO items by their IDs in plain-text format
$ git todo view -q 4e3eeecc 9612977c
4e3eeecc TODO Write some tests as well
9612977c TODO And help people!
```

### Scripting usage examples

For scripting, you should use `-q` or `--quiet` flag to suppress output and avoid cluttering the console.

Also, the `-q`/`--quiet` flag will make the command use a script-friendly output format, which is suitable for parsing in scripts.
However, you still can use `-j`/`--json` or `-p`/`--plain` flags to get the output in JSON or plain text formats, respectively.

```bash
$ git todo view -jq 4e3eeecc | jq '.[] | .title'
"Write some tests as well"
```

# `git todo ls`

This command allows you to print a list of TODO items - optionally, with filtering and various output formats.

## Usage

```bash
git todo ls [-q|--quiet] [-v|--verbose] [-c|--completed] [-i|--incomplete] [-f|--filter <string>] [-j|--json] [-p|--plain]
```

## Flags and arguments

| Flag/Argument                      | Mandatory | Description                                       |
| ---------------------------------- | --------- | ------------------------------------------------- |
| `-q`, `--quiet`                    | No        | Suppress output, only exit codes will be printed. |
| `-v`, `--verbose`                  | No        | Print additional information about the operation. |
| `-c`, `--completed `               | No        | Show only completed TODO items.                   |
| `-i`, `--incomplete `              | No        | Show only incomplete TODO items.                  |
| `-f <string>`, `--filter <string>` | No        | Filter items by a regular expression.             |
| `-j`, `--json`                     | No        | Print TODO items in JSON format.                  |
| `-p`, `--plain`                    | No        | Print TODO items in the plain format.             |

## Filters

There are few built-in filters that can be used to narrow down the list of TODO items:

- `-c` or `--completed`: show only completed TODO items.
  This filter cannot be used together with `-i` or `--incomplete`.
- `-i` or `--incomplete`: show only incomplete TODO items.
  This filter cannot be used together with `-c` or `--completed`.
- `-f <string>` or `--filter <string>`: filter items by a regular expression.
  The filter is case-sensitive and matches the title of the TODO item.

## Output formats

The command supports three output formats:

- **Default**: A human-readable colored format with a numbered list of TODO items:

    ```bash
    $ git todo ls
    e885a108 ✓ Write some code
    4e3eeecc · Write some tests as well
    419ee57f ✓ Write some useful documentation as it is important
    ae19ad18 ✓ Build something great
    9612977c · And help people!
    ```

- **JSON**: A structured JSON format, suitable for parsing in scripts:

    ```bash
    $ git todo ls --json
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
        },
        {
            "id": "419ee57f",
            "done": true,
            "title": "Write some useful documentation as it is important"
        },
        {
            "id": "ae19ad18",
            "done": true,
            "title": "Build something great"
        },
        {
            "id": "9612977c",
            "done": false,
            "title": "And help people!"
        }
    ]
    ```

    The output schema is as follows:

    ```json
    [
        {
            "id"   : <integer>, // The ID of the TODO item
            "done" : <boolean>, // Whether the TODO item is completed (true) or incomplete (false)
            "title": <string>   // The title of the TODO item
        }
    ]
    ```

- **Plain**: A simple text format without any colors or additional formatting:

    ```bash
    $ git todo ls -p
    e885a108 DONE Write some code
    4e3eeecc TODO Write some tests as well
    419ee57f DONE Write some useful documentation as it is important
    ae19ad18 DONE Build something great
    9612977c TODO And help people!
    ```

By default, the command will print the TODO items in a human-readable format.
But if the `-q` or `--quiet` flag is provided, it will force the plain-text format, unless its overriden by any other flags.

## Examples

```bash
# List all TODO items in the default format
$ git todo ls
e885a108 ✓ Write some code
4e3eeecc · Write some tests as well
419ee57f ✓ Write some useful documentation as it is important
ae19ad18 ✓ Build something great
9612977c · And help people!

# List only completed TODO items
$ git todo ls -c
e885a108 ✓ Write some code
419ee57f ✓ Write some useful documentation as it is important
ae19ad18 ✓ Build something great

# List only incomplete TODO items
$ git todo ls -i
4e3eeecc · Write some tests as well
9612977c · And help people!
```

### Scripting usage examples

For scripting, you should use `-q` or `--quiet` flag to suppress output and avoid cluttering the console.

Also, the `-q`/`--quiet` flag will make the command use a script-friendly output format, which is suitable for parsing in scripts.
However, you still can use `-j`/`--json` or `-p`/`--plain` flags to get the output in JSON or plain text formats, respectively.

```bash
# Here is a simple script that shows your progress on TODO items:
$ echo "$(git todo ls -iq | wc -l | xargs)/$(git todo ls -q | wc -l | xargs) TODOs are done"
2/4 TODOs are done
```

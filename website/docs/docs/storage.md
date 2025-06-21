# Storage format

`git-todo` stores its data in a file called `.git/TODO` in the root of your repository.
It's a YAML text file that contains your TODO items in the following format:

```bash
$ cat .git/TODO
items:
    - id: 478e1212
      done: true
      title: Basic CLI
    - id: fa1e7a89
      done: true
      title: Refactor
    - id: 66f3902f
      done: true
      title: Automatable CLI (json output, tab-separated output and other stuff)
    - id: 57ed61f0
      done: true
      title: TUI
    - id: f9eb2530
      done: true
      title: Git hooks
    - id: 69e27356
      title: README
    - id: 39fdec11
      done: true
      title: Tests
    - id: 9e9cf322
      title: Documentation
    - id: "13844228"
      title: CI
    - id: 32fa9037
      title: New TODO Item A
```

You can view and edit this file directly if you prefer,
but it's recommended to use the `git-todo` command-line tool for managing your TODOs.

## Schema

The schema for the `.git/TODO` file is as follows:

```yaml
items:
  - id: <string> # Unique identifier for the TODO item
    done: <boolean> # Whether the TODO item is completed (optional)
    title: <string> # Title of the TODO item
```

| Field           | Type      | Required | Description                                           |
| --------------- | --------- | -------- | ----------------------------------------------------- |
| `items`         | `list`    | No       | List of TODO items                                    |
| `items[].id`    | `string`  | Yes      | Unique identifier for the TODO item                   |
| `items[].done`  | `boolean` | No       | Whether the TODO item is completed (default: `false`) |
| `items[].title` | `string`  | Yes      | Title of the TODO item                                |

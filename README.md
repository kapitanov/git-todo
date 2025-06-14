# git-todo

![GitHub branch check runs](https://img.shields.io/github/check-runs/kapitanov/git-todo/master?label=build)
![GitHub Release](https://img.shields.io/github/v/release/kapitanov/git-todo)
![GitHub License](https://img.shields.io/github/license/kapitanov/git-todo)

A local-only list of TODO items for your git repositories.

---

Are you working on a branch and have some TODO items that you don't want to forget about?

Do you want to be reminded about them before pushing your branch?

Do you want to have a simple way to manage them without using any external services or tools?

If so, `git-todo` is for you!

![TUI](./assets/demo.gif)

## Features

- Keeps a local list of TODO items in your git repository (it's local only and will never be committed and pushed).
- Allows you to add, edit, and remove TODO items.
- Provides a simple text-mode user interface to view and manage TODO items.
- Integrates with git via two hooks:
  - **post-commit**: prints unresolved TODO items after each commit.
  - **pre-push**: prevents pushing branches with unresolved TODO items, but allows manual override.

## How it works

Imagine you are to work on a certain (quite complicated) feature in your git repository - and you would like not to forget about everything you need to implement!

**git-todo** will help you here:

1. First, plan your work and create a list of TODO items - **git-todo** will track them for you.
2. Second, work on your feature, and **git-todo** will remind you about unresolved TODO items before you commit or push your changes.
3. Once you are done with your feature, you can remove the TODO items from the list - and **git-todo** will not bother you anymore.
   
   *Der Mohr hat seine Arbeit getan, der Mohr kann gehen*, as they say.

## Installation

Out of the box, `git-todo` supports macOS, Linux, and Windows.
It is written in Go and can be installed on any platform that supports Go.

We provide pre-built binaries for the most common platforms, so you don't need to build it from sources.
There are several ways to install `git-todo`:

### Install via Homebrew (macOS)

You can install `git-todo` via Homebrew on macOS. This is the recommended way to install it.

```bash
brew tap kapitanov/apps
brew install kapitanov/apps/git-todo
```

### Install from deb package (Ubuntu/Debian)

```bash
export VERSION="0.0.1" # replace with the actual version you want to install
export ARCH="amd64"    # replace with the actual architecture (amd64, arm64, etc.)
wget "https://github.com/kapitanov/git-todo/releases/download/v${VERSION}/git-todo_v${VERSION}_linux_${ARCH}.deb" \
    -O "git-todo_v${VERSION}_linux_${ARCH}.deb"
sudo dpkg -i "git-todo_v${VERSION}_linux_${ARCH}.deb"
git todo --version
```

### Install from rpm package (CentOS/RHEL/Fedora/AWS Linux)

```bash
export VERSION="0.0.1" # replace with the actual version you want to install
export ARCH="amd64"    # replace with the actual architecture (amd64, arm64, etc.)
wget "https://github.com/kapitanov/git-todo/releases/download/v${VERSION}/git-todo_v${VERSION}_linux_${ARCH}.rpm" \
    -O "git-todo_v${VERSION}_linux_${ARCH}.rpm"
sudo rpm -i "git-todo_v${VERSION}_linux_${ARCH}.rpm"
git todo --version
```

### Install from apk package (Alpine)

Run the following commands to install `git-todo` on Alpine Linux:

```bash
export VERSION="0.0.1" # replace with the actual version you want to install
export ARCH="amd64"    # replace with the actual architecture (amd64, arm64, etc.)
wget "https://github.com/kapitanov/git-todo/releases/download/v${VERSION}/git-todo_v${VERSION}_linux_${ARCH}.apk" \
    -O "git-todo_v${VERSION}_linux_${ARCH}.apk"
sudo apk add --allow-untrusted "git-todo_v${VERSION}_linux_${ARCH}.apk"
git todo --version
```

### From Releases (for all platforms)

You can download the latest release of `git-todo` from the [Releases](https://github.com/kapitanov/git-todo/releases) page.
Just pick the version you want and the appropriate binary for your operating system and architecture.

### From sources (for all platforms)

You can install `git-todo` via `go get`:

```bash
go install github.com/kapitanov/git-todo@latest
```

Note that you need to have Go installed on your system to use this method - and you must have the `$GOPATH/bin` in your `PATH`.

## Text-mode User Interface

You can use `git-todo` in a text-mode user interface (TUI) to manage your TODO items
The TUI allows you to navigate through the list of TODO items, toggle their completion status, and exit the interface easily.

```bash
git todo # open a TUI
```

![TUI](./assets/tui.gif)

## Command Line Interface

> TODO: this section will be updated later with more details.

```bash
git todo                                    # open a TUI
git todo init                               # install git hooks
git todo deinit                             # uninstall git hooks
git todo ls                                 # list all TODOs
git todo add "TODO Item"                    # add a new TODO item
git todo add                                # add a new TODO item - opens an editor
git todo check 1                            # check the TODO item with index 1
git todo uncheck 2                          # uncheck the TODO item with index 2
git todo edit 3 --title "Updated TODO Item" # edit the TODO item with index 3
git todo edit 3                             # edit the TODO item with index 3 - opens an editor
git todo rm 4                               # remove the TODO item with index 4
git todo clear                              # clear all TODO items
```

## Storage format

All TODO items are stored locally in `.git/TODO` file in the Markdown-like text format:

```bash
$ cat .git/TODO
[x] Basic CLI
[x] Refactor
[x] Automatable CLI (json output, tab-separated output and other stuff)
[x] TUI
[x] Git hooks
[ ] README
[x] Tests
[ ] Documentation
[ ] CI
[ ] New TODO Item A
```

## Git hooks

Git-todo provides a simple way to manage TODO items in your git repositories without affecting the repository's history or structure. It uses a local `.git/TODO` file to store the items, which is not committed or pushed to the remote repository.

### Installing Git hooks

To install the git hooks, run the following command in your git repository:

```bash
git todo init
```

This will create the necessary hooks in the `.git/hooks` directory.
If there are any existing hooks, they will be backed up to `.git/hooks/<name>.bak` and a new **git-todo** command will be appended to the existing hooks.

It's safe to run this command multiple times, as it will not overwrite existing hooks unless necessary.

### Uninstalling Git hooks

To uninstall the git hooks, run the following command in your git repository:

```bash
git todo deinit
```

This will remove the git-todo command from the hooks. Again, if there are any existing hooks, they will be backed up to `.git/hooks/<name>.bak`.

### Pre-commit hook

The pre-commit hook is a simple way to remind you about unresolved TODO items before committing your changes.
It will print a warning message if there are any unresolved items and will prevent the commit(unless you confirm it by typing `y`):

```bash
$ git commit -m "Commit message"
You still have some TODO items to resolve:
 - README
 - Documentation
 - CI
 - New TODO Item A

Are you sure you want to commit these changes to the "master" branch (y/n)?
```

Unless you type `y`, the commit will be aborted and you will be returned to the command line.

### Pre-push hook

The pre-push hook is a simple way to remind you about unresolved TODO items before pushing your changes.
It will print a warning message if there are any unresolved items and will prevent the push unless you confirm it by typing `y`:

```bash
$ git push origin my-branch-name
You still have some TODO items to resolve:
 - README
 - Documentation
 - CI
 - New TODO Item A
Are you sure you want to push these changes (y/n)?
```

Unless you type `y`, the push will be aborted and you will be returned to the command line.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

# Installation Guide

Out of the box, `git-todo` supports macOS, Linux, and Windows.
It is written in Go and can be installed on any platform that supports Go.

We provide pre-built binaries for the most common platforms, so you don't need to build it from sources.
There are several ways to install `git-todo`:

- Install via Homebrew
- Install from Linix packages (deb/rpm/apk)
- Install from Github Releases
- Install from sources

### Install via Homebrew

You can install `git-todo` via Homebrew on macOS. This is the recommended way to install it.

```shell
brew tap kapitanov/apps
brew install kapitanov/apps/git-todo
```

Copy these commands into your terminal and `git-todo` will be installed momentarily using Homebrew.

### Install from Linix packages

We provide pre-built packages for the most common Linux distributions.
You can install `git-todo` using the package manager of your distribution.

=== "Ubuntu/Debian"

    ```bash
    export VERSION="0.0.1" # replace with the actual version you want to install
    export ARCH="amd64"    # replace with the actual architecture (amd64, arm64, etc.)
    wget "https://github.com/kapitanov/git-todo/releases/download/v${VERSION}/git-todo_v${VERSION}_linux_${ARCH}.deb" \
        -O "git-todo_v${VERSION}_linux_${ARCH}.deb"
    sudo dpkg -i "git-todo_v${VERSION}_linux_${ARCH}.deb"
    git todo --version
    ```

=== "CentOS/RHEL/Fedora/AWS Linux"

    ```bash
    export VERSION="0.0.1" # replace with the actual version you want to install
    export ARCH="amd64"    # replace with the actual architecture (amd64, arm64, etc.)
    wget "https://github.com/kapitanov/git-todo/releases/download/v${VERSION}/git-todo_v${VERSION}_linux_${ARCH}.rpm" \
        -O "git-todo_v${VERSION}_linux_${ARCH}.rpm"
    sudo rpm -i "git-todo_v${VERSION}_linux_${ARCH}.rpm"
    git todo --version
    ```

=== "Alpine"

    ```bash
    export VERSION="0.0.1" # replace with the actual version you want to install
    export ARCH="amd64"    # replace with the actual architecture (amd64, arm64, etc.)
    wget "https://github.com/kapitanov/git-todo/releases/download/v${VERSION}/git-todo_v${VERSION}_linux_${ARCH}.apk" \
        -O "git-todo_v${VERSION}_linux_${ARCH}.apk"
    sudo apk add --allow-untrusted "git-todo_v${VERSION}_linux_${ARCH}.apk"
    git todo --version
    ```

Copy these commands into your terminal and `git-todo` will be installed momentarily using your package manager.

!!! note

    You need to replace the `VERSION` and `ARCH` variables with the actual version and architecture you want to install.
    You can find the latest version on our [Releases](https://github.com/kapitanov/git-todo/releases) page,
    and the architecture can be `amd64`, `arm`, or `arm64` depending on your system.

### Install from Github Releases

You can download the latest release of `git-todo` from the [Releases](https://github.com/kapitanov/git-todo/releases) page.
Just pick the version you want and the appropriate binary for your operating system and architecture.

We provide various package formats: not only `deb`, `rpm`, and `apk`,
but also standalone binaries for Windows, macOS, and Linux (packed into a zip or tar.gz-archive).
SHA256 checksums are provided for all packages, so you can verify the integrity of the downloaded files.

### Install from sources

And after all, if you want to build `git-todo` from sources, you can do it easily:

=== "Via `go install`"

    1. First, install Go on your system if you haven't done it yet.
    2. Then, run the following command to fetch the repository, build the project and install its binary file into your `$GOPATH/bin` directory:

       ```bash
       go install github.com/kapitanov/git-todo@latest
       ```

=== "Via `git clone`"

    1. First, install Go on your system if you haven't done it yet.
    2. Then, run the following command to fetch the repository, build the project and install its binary file into your `$GOPATH/bin` directory:

       ```bash
       git clone git@github.com:kapitanov/git-todo.git
       cd git-todo
       make install
       ```

!!! note

    You need to have Go installed on your system to use this method - and you must have the `$GOPATH/bin` in your `PATH`.

# Quickstart Guide

Have you ever been working on a project - say, on a new feature - and rushed into commiting your changes,
only to realize later that you forgot to address some important things? 
Perhaps you even had a list of tasks to complete, but you forgot to check it before committing?

If this sounds familiar, then `git-todo` is the tool for you!

`git-todo` is a command-line tool that helps you manage your TODO items in Git repositories.
It allows you to create, view, and manage TODO items directly from the command line,
ensuring that you never forget to address important tasks before committing your changes.
And its [git hooks](./git-hooks.md) will help you to automate the management of TODO items,
ensuring that you always address your tasks before committing or pushing changes.

All your TODO items are stored locally in your Git repository,
so you can work offline and still keep track of your tasks.
And they will never be committed and pushed to remote repositories: your TODOs are for your eyes only!
Remember: `git-todo` is not a task management tool, it is a tool to help you keep your work on the code a little bit more organized.

Let's get started!

1. Install `git-todo` by following the [installation instructions](./install.md).
2. Open a terminal and navigate to your Git repository where you want to use `git-todo`.
   If you don't have a Git repository yet, you can create one by running:

    ```bash
    $ git init my-repo
    $ cd my-repo
    ```

3. Initialize `git-todo` in your Git repository by running the following command:

    ```bash
    $ git todo init
    ```

    This will set up the necessary hooks and create a `.git-todo` directory in your repository.

4. Now, plan your work! Think about the tasks you need to complete and create a TODO item for each task:

    ```bash
    $ git todo add "Write some code"
    Added new TODO item: "Write some code"

    $ git todo add "Write tests for the code"
    Added new TODO item: "Write tests for the code"

    $ git todo add "Update documentation"
    Added new TODO item: "Update documentation"

    $ git todo add "Check if everything works"
    Added new TODO item: "Check if everything works"
    ```

5. Now you may work on your tasks.
   When you finish a task, you can mark it as done:

    ```bash
    $ git todo ls
    1 · Write some code
    2 · Write tests for the code
    3 · Update documentation
    4 · Check if everything works

    $ git todo check 1
    TODO item #1 has been checked as completed (Write some code)

    $ git todo ls
    1 ✓ Write some code
    2 · Write tests for the code
    3 · Update documentation
    4 · Check if everything works
    ```

6. Once you have completed all your tasks, you can commit your changes:

    ```bash
    $ git add .
    $ git commit -m "Completed tasks"
    ```

    But if you try to commit (or push) while there are still incomplete TODO items, `git-todo` will prevent you from doing so:

    ```bash
    $ git commit -m "Incomplete tasks"
    You still have some TODO items to resolve:
     - Write tests for the code
     - Update documentation
     - Check if everything works
    Are you sure you want to commit these changes to the "fix-bug-in-login" branch (y/n) ? n
    commit aborted due to incomplete TODO items
    ```

# Git Hooks

While the first half of `git-todo`'s core functionality is focused on managing TODO items,
the second half is about integrating these items into your Git workflow through hooks.

!!! note

    Git hooks are scripts that run automatically at certain points in the Git workflow.
    For more information about them, please refer to the [Git hooks documentation](https://git-scm.com/book/en/v2/Customizing-Git-Git-Hooks).

`git-todo` provides several hooks that can be used to automate the management of TODO items and help you keep track of your tasks as you work with Git:

- **pre-commit**: This hook runs before a commit is made.
  It can be used to check for TODO items that need attention before committing changes.
- **pre-push**: This hook runs before pushing changes to a remote repository.
  It can be used to ensure that all TODO items are addressed before pushing.

## Installing

Initially, no hooks are installed, and you need to set them up manually.
To install the hooks, you can use the `git-todo` command:

```bash
$ git todo init
```

This command will gracefully install the necessary hooks in your current Git repository.
For more information about the `git-todo init` command, see the [corresponding documentation](./cli/init.md).

## Pre-commit Hook

The pre-commit hook is designed to run before a commit is made.
It checks for TODO items that need attention and prompts you to address them before proceeding with the commit.
If there are any TODO items that are not completed, the hook will prevent the commit from being made
and display a message with the list of incomplete TODO items.
You still can proceed with the commit by typing `y` or `yes` when prompted.

Here is an example of how the pre-commit hook works:

```bash
$ git commit -m "My commit message"
You still have some TODO items to resolve:
 - Fix the bug in the login feature
 - Update the documentation
 - Refactor the code for better readability
Are you sure you want to commit these changes to the "fix-bug-in-login" branch (y/n) ? n
commit aborted due to incomplete TODO items

$ git commit -m "My commit message"
You still have some TODO items to resolve:
 - Fix the bug in the login feature
 - Update the documentation
 - Refactor the code for better readability
Are you sure you want to commit these changes to the "fix-bug-in-login" branch (y/n) ? y
[fix-bug-in-login 021425b] My commit message
 1 file changed, 1 insertion(+)
 create mode 100644 bugfix.txt
```

This hook will have no effect if there are no unresolved TODO items in the current git repository.

## Pre-push Hook

The pre-push hook is designed to run before pushing changes to a remote repository.
It checks for TODO items that need attention and prompts you to address them before proceeding with the push.
If there are any TODO items that are not completed, the hook will prevent the push from being made
and display a message with the list of incomplete TODO items.
You still can proceed with the push by typing `y` or `yes` when prompted.

```bash
git push
You still have some TODO items to resolve:
 - Fix the bug in the login feature
 - Update the documentation
 - Refactor the code for better readability
Are you sure you want to push these changes (y/n) ? n
push aborted due to incomplete TODO items

git push
You still have some TODO items to resolve:
 - Fix the bug in the login feature
 - Update the documentation
 - Refactor the code for better readability
Are you sure you want to push these changes (y/n) ? y
Enumerating objects: 4, done.
Counting objects: 100% (4/4), done.
Delta compression using up to 8 threads
Compressing objects: 100% (2/2), done.
Writing objects: 100% (3/3), 331 bytes | 331.00 KiB/s, done.
Total 3 (delta 0), reused 0 (delta 0), pack-reused 0 (from 0)
To git://github.com/yourusername/yourrepository.git
   021425b..427235b  fix-bug-in-login -> fix-bug-in-login
```

This hook will have no effect if there are no unresolved TODO items in the current git repository.

## Uninstalling

To uninstall the hooks, you can use the `git-todo` command:

```bash
$ git todo deinit
```

This command will gracefully remove the necessary hooks from your current Git repository.
For more information about the `git-todo deinit` command, see the [corresponding documentation](./cli/deinit.md).

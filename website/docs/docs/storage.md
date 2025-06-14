# Storage format

`git-todo` stores its data in a file called `.git/TODO` in the root of your repository.
It's a simple plain text file that contains your TODO items, each on a separate line,
in the following format:

```
[x] Write some code
[ ] Write tests for the code
[ ] Update documentation
[ ] Check if everything works
```

Where `[x]` indicates a completed TODO item and `[ ]` indicates an incomplete one.
You can view and edit this file directly if you prefer,
but it's recommended to use the `git-todo` command-line tool for managing your TODOs.

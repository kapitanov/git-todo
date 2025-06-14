# Interaction with IDEs

The major downside of the current version of the `git-todo` tool is that it does not integrate with IDEs.
This means not only that you cannot see your TODOs in your IDE, and you cannot edit them from your IDE,
but (more importantly) you cannot use the IDE's git-related features like committing and pushing changes,
as the git hooks will prevent IDE from any actions.

<figure markdown="span">
  ![](../assets/vscode.png){ width="300" }
  <figcaption>
    Here is an example: that's what happens when you try to commit changes in VSCode while having some unresolved TODOs.
  </figcaption>
</figure>

Still, `git-todo` will still fulfill its purpose, as it will prevent you from committing changes that are not ready to be committed.
You would just have to keep in mind that you will have to resolve all TODOs before committing changes.

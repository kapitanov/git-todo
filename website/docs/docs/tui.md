# Text-mode User Interface (TUI)

While `git-todo` is primarily a command-line tool,
it also provides a Text-mode User Interface (TUI) for users who prefer a more interactive experience.
The TUI allows you to navigate through your TODO items and manage them in a user-friendly manner.

<figure markdown="span">
  ![](../assets/tui.gif){ width="800px" }
</figure>

To open the TUI, simply run:

```bash
$ git todo
```

This will launch the TUI, where you can view, add, edit, and delete TODO items using keyboard shortcuts:

| Shortcut                       | Action                             |
| ------------------------------ | ---------------------------------- |
| ++up++ or ++k++                | Move selection up                  |
| ++down++ or ++j++              | Move selection down                |
| ++n++                          | Open new TODO item creation dialog |
| ++space++ or ++t++             | Toggle completion status           |
| ++d++                          | Open TODO item deletion dialog     |
| ++e++ or ++enter++             | Open TODO item editing dialog      |
| ++x++                          | Open TODO items clear dialog       |
| ++q++ or ++esc++ or ++ctrl+c++ | Quit the TUI                       |

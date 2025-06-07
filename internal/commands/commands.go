package commands

import (
	"fmt"
	"io"
	"runtime/debug"
	"strconv"

	"github.com/spf13/cobra"

	"github.com/kapitanov/git-todo/internal/application"
	"github.com/kapitanov/git-todo/internal/commands/tui"
	"github.com/kapitanov/git-todo/internal/logutil"
)

func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "git-todo",
		Short: "A simple todo list for git",
		Long: `git-todo is a simple todo list for git. Manage your tasks with ease using this command-line tool.
It allows you to manage a private, working-copy-local TODO list.

Features:
  - Manage private, per-repository TODO items
  - Use a Text User Interface (TUI) to interact with your TODO items
  - Convinient git hooks will prevent you from accidentally pushing your changes with uncompleted TODO items
    or committing them into the main branch

Quickstart:
  1. cd to your git repository
  2. Run 'git-todo init' to set up git todo hooks
  3. Run 'git-todo' to open the Text User Interface (TUI) for git todo
  4. Work on your tasks - and when you're done, commit your changes as usual.
  5. If you try to push your changes with uncompleted TODO items, git todo will warn you.
`,
		Example: `  git-todo                                              - open the Text User Interface (TUI) for git todo
  git-todo init                                         - initialize git todo hooks for the current repository
  git-todo add Implement new feature                    - add a TODO item with the description "Implement new feature"
  git-todo add "Write some documentation"               - add another TODO item
  git-todo ls                                           - list all TODO items
  git-todo edit 1 --title "Update task description"     - edit the first TODO item, changing its title
  git-todo check 1                                      - mark the first TODO item as completed
  git-todo uncheck 1                                    - mark the first TODO item as not completed
  git-todo remove 2                                     - remove the second TODO item
  git-todo clear                                        - remove all TODO items
`,
		SilenceUsage:  true,
		Version:       getAppVersion(),
		SilenceErrors: true,
		CompletionOptions: cobra.CompletionOptions{
			HiddenDefaultCmd: true,
		},
	}

	var verbose, quiet bool
	cmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "enable verbose output")
	cmd.PersistentFlags().BoolVarP(&quiet, "quiet", "q", false, "suppress all unnecessary output")

	cmd.PersistentPreRunE = func(cmd *cobra.Command, args []string) error {
		if quiet {
			verbose = false
			cmd.SetErr(io.Discard)
		}

		logutil.ConfigureLogger(verbose)
		return nil
	}

	cmd.RunE = func(cmd *cobra.Command, _ []string) error {
		app, err := application.New()
		if err != nil {
			return err
		}

		err = tui.Run(cmd.Context(), app)
		if err != nil {
			return err
		}
		return nil
	}

	cmd.AddCommand(List())
	cmd.AddCommand(Add())
	cmd.AddCommand(Edit())
	cmd.AddCommand(Check())
	cmd.AddCommand(Uncheck())
	cmd.AddCommand(Remove())
	cmd.AddCommand(Clear())
	cmd.AddCommand(Path())
	cmd.AddCommand(Init())
	cmd.AddCommand(Deinit())
	cmd.AddCommand(GitHooks())

	return cmd
}

func getAppVersion() string {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		return "unknown"
	}

	return info.Main.Version
}

func selectItemsByIndex(app *application.App, args []string) (items []*application.Item, err error) {
	for _, arg := range args {
		index, err := strconv.Atoi(arg)
		if err != nil {
			return nil, fmt.Errorf("invalid index %q", arg)
		}

		item := app.Item(index)
		if item == nil {
			return nil, fmt.Errorf("item #%d doesn't exist", index)
		}

		items = append(items, item)
	}
	return
}

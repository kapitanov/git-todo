package commands

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/kapitanov/git-todo/internal/git"
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
		SilenceErrors: true,
		CompletionOptions: cobra.CompletionOptions{
			HiddenDefaultCmd: true,
		},
	}

	commandCtx := &commandContext{}
	cmd.PersistentFlags().BoolVarP(&commandCtx.IsVerbose, "verbose", "v", false, "enable verbose output")
	cmd.PersistentFlags().BoolVarP(&commandCtx.IsQuiet, "quiet", "q", false, "suppress all unnecessary output")

	cmd.PersistentPreRunE = func(cmd *cobra.Command, args []string) error {
		if commandCtx.IsQuiet {
			commandCtx.IsVerbose = false
			cmd.SetErr(io.Discard)
		}

		logutil.ConfigureLogger(commandCtx.IsVerbose)
		return nil
	}

	cmd.RunE = func(cmd *cobra.Command, _ []string) error {
		app, err := application.New()
		if err != nil {
			return err
		}

		err = tui.Run(cmd.Context(), app)
		if err != nil {
			return commandCtx.HandleError(err)
		}
		return nil
	}

	cmd.AddCommand(listCommand(commandCtx))
	cmd.AddCommand(addCommand(commandCtx))
	cmd.AddCommand(editCommand(commandCtx))
	cmd.AddCommand(checkCommand(commandCtx))
	cmd.AddCommand(uncheckCommand(commandCtx))
	cmd.AddCommand(removeCommand(commandCtx))
	cmd.AddCommand(clearCommand(commandCtx))
	cmd.AddCommand(pathCommand(commandCtx))
	cmd.AddCommand(initCommand(commandCtx))
	cmd.AddCommand(deinitCommand(commandCtx))
	cmd.AddCommand(gitHooksCommand(commandCtx))

	return cmd
}

func selectItemsByIndex(app *application.App, args []string) (items []*application.Item, err error) {
	for _, arg := range args {
		index, err := strconv.Atoi(arg)
		if err != nil {
			return nil, fmt.Errorf("invalid index %q", arg)
		}

		item := app.Item(index)
		if item == nil {
			return nil, ExitError{
				ExitCode: ExitCodeItemDoesntExist,
				Message:  fmt.Sprintf("item %d doesn't exist", index),
			}
		}

		items = append(items, item)
	}
	return
}

type ExitCode int

const (
	ExitCodeNotGitRepository  ExitCode = 128
	ExitCodeItemDoesntExist   ExitCode = 404
	ExitCodeItemAlreadyExists ExitCode = 409
	ExitCodeOperationCanceled ExitCode = 499
	ExitCodeInternalError     ExitCode = 500
)

type ExitError struct {
	ExitCode ExitCode
	Message  string
}

func (e ExitError) Error() string {
	if e.Message != "" {
		return e.Message
	}

	return fmt.Sprintf("exit code %d", e.ExitCode)
}

type commandContext struct {
	IsVerbose bool
	IsQuiet   bool
}

func (c *commandContext) IsRunningInInteractiveMode() bool { return !c.IsQuiet }

func (c *commandContext) HumanReadablePrintf(format string, args ...any) {
	if !c.IsRunningInInteractiveMode() {
		return
	}

	_, _ = fmt.Fprintf(os.Stdout, format, args...)
}

func (c *commandContext) MachineReadablePrintf(format string, args ...any) {
	if c.IsRunningInInteractiveMode() {
		return
	}

	_, _ = fmt.Fprintf(os.Stdout, format, args...)
}

func (c *commandContext) HandleError(err error) error {
	if errors.Is(err, git.ErrNoGitRepository) {
		_, _ = fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		return ExitError{ExitCode: ExitCodeNotGitRepository}
	}

	if errors.Is(err, application.ErrItemAlreadyExists) {
		_, _ = fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		return ExitError{ExitCode: ExitCodeItemAlreadyExists}
	}

	if errors.Is(err, context.Canceled) {
		c.HumanReadablePrintf("%s\n", "Cancelled by user")
		return ExitError{ExitCode: ExitCodeOperationCanceled}
	}
	return err
}

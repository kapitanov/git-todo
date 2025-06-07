package commands

import (
	"strings"

	"github.com/spf13/cobra"

	"github.com/kapitanov/git-todo/internal/application"
	"github.com/kapitanov/git-todo/internal/commands/cui"
)

func Add() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add [TITLE]...",
		Short: "add a new TODO item",
		Long: `Add a new TODO item to your list. You can provide the title as arguments or edit it interactively.
If any arguments are provided, they will be treated as the title of the new TODO item.
Otherwise, an interactive editor will be opened to type the title.

By default, the editor is determined by the EDITOR environment variable.
If it is not set, it will fall back to the system's default editor (usually vim or nano).
You can override this by setting the EDITOR environment variable before running the command.`,
		Example: `  git todo add "Write some code"        - adds a new incomplete TODO item with the title "Write some code"
  git todo add Write some tests as well - adds a new incomplete TODO item with the title "Write some tests as well"
  git todo add                          - opens an interactive editor to type the title of the new TODO item
  EDITOR=nano git todo add              - same as above, but uses nano as the editor`,
		Args: cobra.ArbitraryArgs,
	}

	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		app, err := application.New()
		if err != nil {
			return err
		}

		title := strings.Join(args, " ")

		if title == "" {
			title, err = cui.Edit("", "Type the title of the new TODO item")
			if err != nil {
				return err
			}

			if title == "" {
				return nil
			}
		}

		_, err = app.NewItem(title)
		if err != nil {
			return err
		}

		cmd.PrintErrf("Added new TODO item: %q\n", title)
		return nil
	}

	return cmd
}

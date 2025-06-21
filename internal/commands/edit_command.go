package commands

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/kapitanov/git-todo/internal/application"
	"github.com/kapitanov/git-todo/internal/commands/cui"
)

func editCommand(c *commandContext) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "edit <id>",
		Short: "edit a TODO item",
		Long: `Edit a TODO item by its ID.
You may find the IDs using the "git todo ls" command.

This command allows you to change the title of a TODO item.
If any arguments are provided, they will be treated as the new title of the TODO item.
Otherwise, an interactive editor will be opened to type the title.

By default, the editor is determined by the EDITOR environment variable.
If it is not set, it will fall back to the system's default editor (usually vim or nano).
You can override this by setting the EDITOR environment variable before running the command.`,
		Example: `  git todo edit 4e3eeecc -t "New title" - edit TODO item [4e3eeecc] and set its title to "New title"
  git todo edit 4e3eeecc                - opens an interactive editor to type the new title of the TODO item [4e3eeecc]
  EDITOR=nano git todo edit 4e3eeecc    - same as above, but uses nano as the editor
  `,
		Args: cobra.ExactArgs(1),
	}

	var title string

	cmd.Flags().StringVarP(&title, "title", "t", "", "new title of the TODO item")

	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		app, err := application.New()
		if err != nil {
			return c.HandleError(err)
		}

		id := args[0]
		item := app.Item(id)
		if item == nil {
			err = ExitError{
				ExitCode: ExitCodeItemDoesntExist,
				Message:  fmt.Sprintf("item [%s] doesn't exist", id),
			}
			return c.HandleError(err)
		}

		if title == "" {
			if !c.IsRunningInInteractiveMode() {
				return errors.New("no title provided")
			}

			hint := fmt.Sprintf("Type the new title of TODO item [%s]\n Old title: %q\n", item.ID(), item.Title())
			title, err = cui.Edit(item.Title(), hint)
			if err != nil {
				return err
			}
		}

		if title == "" {
			return nil
		}

		originalTitle := item.Title()
		err = item.SetTitle(title)
		if err != nil {
			return c.HandleError(err)
		}

		c.MachineReadablePrintf("%s\n", item.ID())
		if originalTitle != title {
			c.HumanReadablePrintf("TODO item [%s] has been renamed:\n  old: %q\n  new: %q\n", item.ID(), originalTitle, title)
		} else {
			c.HumanReadablePrintf("TODO item [%s] has not been renamed: the new title is the same as the old one\n", item.ID())
		}
		return nil
	}

	return cmd
}

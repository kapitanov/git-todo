package commands

import (
	"errors"
	"strings"

	"github.com/spf13/cobra"

	"github.com/kapitanov/git-todo/internal/application"
	"github.com/kapitanov/git-todo/internal/commands/cui"
)

func addCommand(c *commandContext) *cobra.Command {
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
  EDITOR=nano git todo add              - same as above, but uses nano as the editor
  git todo add -u "Write some code"     - adds a new incomplete TODO item with the title "Write some code",
                                          or returns an existing item in case of conflict`,
		Args: cobra.ArbitraryArgs,
	}

	var skipExisting bool
	cmd.Flags().BoolVarP(&skipExisting, "unless-exists", "u", false, "skip adding the item if it already exists in the list")

	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		app, err := application.New()
		if err != nil {
			return c.HandleError(err)
		}

		title := strings.Join(args, " ")

		if title == "" {
			if !c.IsRunningInInteractiveMode() {
				return errors.New("no title provided")
			}

			title, err = cui.Edit("", "Type the title of a new TODO item")
			if err != nil {
				return c.HandleError(err)
			}

			if title == "" {
				return nil
			}
		}

		newItemAdded := true
		item, err := app.NewItem(title)
		if err != nil {
			var errAlreadyExists application.ItemAlreadyExistsError
			if skipExisting && errors.As(err, &errAlreadyExists) {
				item = errAlreadyExists.Item
				newItemAdded = false
			} else {
				return c.HandleError(err)
			}
		}

		c.MachineReadablePrintf("%s\n", item.ID())
		if newItemAdded {
			c.HumanReadablePrintf("Added a new TODO item: [%s] %q\n", item.ID(), item.Title())
		} else {
			c.HumanReadablePrintf("The TODO item already exists: [%s] %q\n", item.ID(), item.Title())
		}
		return nil
	}

	return cmd
}

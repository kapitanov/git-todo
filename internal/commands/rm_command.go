package commands

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/kapitanov/git-todo/internal/application"
	"github.com/kapitanov/git-todo/internal/commands/cui"
)

func removeCommand(c *commandContext) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "rm",
		Short: "remove a TODO item",
		Long: `Remove a TODO item by its index in the list.
You may find the indices using the "git todo ls" command.

This action cannot be undone!

This command requires a confirmation before removing the item.
If you want to skip the confirmation, use the --force flag.
`,
		Example: `  git todo rm 5        - removes TODO item with index 5
  git todo rm 1 3 5    - removes TODO items with indices 1, 3, and 5
  git todo rm --force5 - removes TODO items with index 5 without confirmation`,
		Args: cobra.MinimumNArgs(1),
	}

	var force bool
	cmd.Flags().BoolVarP(&force, "force", "f", false, "force remove the TODO item without confirmation")

	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		if !force && !c.IsRunningInInteractiveMode() {
			return errors.New("this command requires either an interactive confirmation prompt or a \"--force\" flag")
		}

		app, err := application.New()
		if err != nil {
			return c.HandleError(err)
		}

		items, err := selectItemsByIndex(app, args)
		if err != nil {
			return c.HandleError(err)
		}

		for _, item := range items {
			if !force {
				confirmed, err := cui.Confirm(fmt.Sprintf("Are you sure you want to remove TODO item #%d %q", item.ID(), item.Title()))
				if err != nil {
					return c.HandleError(err)
				}
				if !confirmed {
					c.HumanReadablePrintf("Canceled removal of TODO item %d %q\n", item.ID(), item.Title())
					continue
				}
			}

			err = item.Delete()
			if err != nil {
				return c.HandleError(err)
			}

			c.HumanReadablePrintf("TODO item %d has been removed (%s)\n", item.ID(), item.Title())
			c.MachineReadablePrintf("%d\n", item.ID())
		}
		return nil
	}

	return cmd
}

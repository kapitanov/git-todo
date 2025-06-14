package commands

import (
	"github.com/spf13/cobra"

	"github.com/kapitanov/git-todo/internal/application"
)

func checkCommand(c *commandContext) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "check <index>...",
		Short: "mark a TODO item as completed",
		Long: `Mark one or more TODO items as completed by their index in the list.
You may find the indices using the "git todo ls" command.

If an item is already marked as completed, this command won't do anything.`,
		Example: `  git todo check 5     - marks TODO item with index 5 as completed
  git todo check 1 3 5 - marks TODO items with indices 1, 3, and 5 as completed`,
		Args: cobra.MinimumNArgs(1),
	}

	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		app, err := application.New()
		if err != nil {
			return c.HandleError(err)
		}

		items, err := selectItemsByIndex(app, args)
		if err != nil {
			return c.HandleError(err)
		}

		for _, item := range items {
			if item.IsCompleted() {
				c.HumanReadablePrintf("TODO item %d is already marked as completed (%s)\n", item.ID(), item.Title())
				continue
			}

			err = item.SetIsCompleted(true)
			if err != nil {
				return c.HandleError(err)
			}

			c.HumanReadablePrintf("TODO item %d has been marked as completed (%s)\n", item.ID(), item.Title())
			c.MachineReadablePrintf("%d\n", item.ID())
		}
		return nil
	}

	return cmd
}
